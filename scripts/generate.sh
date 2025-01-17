#!/bin/bash -e

TEMPDIR=$(mktemp -d)
mkdir -p $TEMPDIR/src/google.golang.org
ln -s $PWD/_output/src/google.golang.org/protobuf $TEMPDIR/src/google.golang.org/protobuf
function cleanup {
    # echo $TEMPDIR
    unlink $TEMPDIR/src/google.golang.org/protobuf
    rm -rf $TEMPDIR
}
trap cleanup EXIT

# Ensure we're using protoc and gomvpkg that this repo downloads.
export PATH=$PWD/_output/bin:$PATH

# Copy all .proto files from Kubernetes into a temporary directory.
REPOS=( "apimachinery" "api" "apiextensions-apiserver" "kube-aggregator" )
for REPO in "${REPOS[@]}"; do
    SOURCE=$PWD/_output/kubernetes/staging/src/k8s.io/$REPO
    TARGET=$TEMPDIR/src/k8s.io
    mkdir -p $TARGET
    rsync -a --prune-empty-dirs --include '*/' --include '*.proto' --exclude '*' $SOURCE $TARGET
done

function fix_option {
    filename=$1
    dirname=${filename%/*}
    pkgname="k8s.io/${dirname#*k8s.io/}"
    vername=${pkgname##*/}
    script="/option go_package/s#${vername}#${pkgname}#"
    sed -i '' "${script}" $1
}

export -f fix_option
find $TEMPDIR/src/k8s.io -name '*.proto' -exec bash -c 'fix_option "${1}"' -- {} \;

# Remove API groups that aren't actually real
rm -r $TEMPDIR/src/k8s.io/apimachinery/pkg/apis/testapigroup

cd $TEMPDIR/src
for FILE in $( find . -type f ); do
    protoc --go_out=. $FILE
done
grep -r -h go_package $TEMPDIR/src
rm $( find . -type f -name '*.proto' );
cd -

export GOPATH=$TEMPDIR
function mvpkg {
    FROM="k8s.io/$1"
    TO="github.com/onedata/k8s/$2"
    mkdir -p "$GOPATH/src/$(dirname $TO)"
    echo "gompvpkg -from=$FROM -to=$TO"
    gomvpkg -from=$FROM -to=$TO
}

mvpkg apiextensions-apiserver/pkg/apis/apiextensions/v1beta1 apis/apiextensions/v1beta1
mvpkg apiextensions-apiserver/pkg/apis/apiextensions/v1 apis/apiextensions/v1
mvpkg apimachinery/pkg/api/resource apis/resource
mvpkg apimachinery/pkg/apis/meta apis/meta
mvpkg apimachinery/pkg/runtime runtime
mvpkg apimachinery/pkg/util util
for DIR in $( ls ${TEMPDIR}/src/k8s.io/api/ ); do
    mvpkg api/$DIR apis/$DIR
done
mvpkg kube-aggregator/pkg/apis/apiregistration apis/apiregistration

rm -rf api apis runtime util
mv $TEMPDIR/src/github.com/onedata/k8s/* .
