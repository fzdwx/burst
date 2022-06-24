VERSION=1.5
ACTION=$1

function build_server() {
    echo "start build server jar"
    mvn -f ../pom.xml clean package
    echo "start build server docker image"
    cd `dirname $0`
    CURRENT_DIR=`pwd`
    echo $CURRENT_DIR
    cd ../burst-server/target
    CONTEXT_DIR=`pwd`
    echo $CONTEXT_DIR
    nohup java -jar burst-server-$VERSION.jar > /root/burst-server.log 2>&1 &
#    echo "docker build -t fzdwx/burst-server:$VERSION -f $CURRENT_DIR/Dockerfile $CONTEXT_DIR"
#    docker build -t fzdwx/burst-server:$VERSION -f $CURRENT_DIR/Dockerfile $CONTEXT_DIR
}

case "$ACTION" in
server)
    build_server
  ;;
esac