VERSION=1.3
ACTION=$1

function build_server() {
    cd `dirname $0`
    CURRENT_DIR=`pwd`
    cd ../burst-server/target
    CONTEXT_DIR=`pwd`
    echo "docker build -t fzdwx/burst-server:$VERSION -f $CURRENT_DIR/Dockerfile $CONTEXT_DIR"
}

case "$ACTION$" in
server)
    build_server
  ;;
esac