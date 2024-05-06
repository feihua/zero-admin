#!/bin/bash

APP_NAME=$1

function start()
{
    count=`ps -ef grep $APP_NAME|grep -v grep|wc -l`
    if [ $count != 0 ];then
        echo "$APP_NAME is running..."
    else
        nohup ./"$APP_NAME" -f ./"$APP_NAME".yaml  > /dev/null 2>&1 &
        echo -e "Start $APP_NAME success...Please see the detail log in logs"
    fi
}

function stop()
{
    echo "Stop $APP_NAME"
    pid=`ps -ef |grep java|grep $APP_NAME|grep -v grep|awk '{print $2}'`
    count=`ps -ef |grep $APP_NAME|grep -v grep|wc -l`

    if [ $count != 0 ];then
        kill $pid
        count=`ps -ef |grep $APP_NAME|grep -v grep|wc -l`

        pid=`ps -ef |grep $APP_NAME|grep -v grep|awk '{print $2}'`
        kill -9 $pid
    fi
}

function restart()
{
    stop
    sleep 2
    start
}

function status()
{
    count=`ps -ef |grep $APP_NAME|grep -v grep|wc -l`
    if [ $count != 0 ];then
        echo "$APP_NAME is running..."
    else
        echo "$APP_NAME is not running..."
    fi
}

case $2 in
    start)
    start;;
    stop)
    stop;;
    restart)
    restart;;
    status)
    status;;
    *)

    echo -e "\033[0;31m Usage: \033[0m  \033[0;34m sh  $0  {start|stop|restart|status}  {JenkinsJarName} \033[0m
\033[0;31m Example: \033[0m
      \033[0;33m sh  $0  start esmart-test.jar \033[0m"
esac

