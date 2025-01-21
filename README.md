```shell
pipeline {
agent any

    stages {
        stage('clone') {
            steps {
                git 'https://gitee.com/liufeihua/zero-admin.git'
            }
        }
        stage('install') {
            steps {
                    sh '''
                   export GOPROXY=https://goproxy.cn,direct
                   /usr/local/go/bin/go mod tidy
                    '''
            }
        }
        stage('build service') {
          parallel {
            stage('sys-rpc') {
                steps {
                    sh '''
                    mkdir -p target/sys-rpc
                    cp rpc/sys/etc/sys.yaml target/sys-rpc/sys-rpc.yaml
                    /usr/local/go/bin/go build -o target/sys-rpc/sys-rpc -v ./rpc/sys/sys.go
                    '''
                }
            }
            stage('ums-rpc') {
                steps {
                    sh '''
                    mkdir -p target/ums-rpc
                    cp rpc/ums/etc/ums.yaml target/ums-rpc/ums-rpc.yaml
                    /usr/local/go/bin/go build -o target/ums-rpc/ums-rpc -v ./rpc/ums/ums.go
                    '''
                     
                }
    
            }
            stage('oms-rpc') {
                steps {
                    sh '''
                    mkdir -p target/oms-rpc
                    cp rpc/oms/etc/oms.yaml target/oms-rpc/oms-rpc.yaml
                    /usr/local/go/bin/go build -o target/oms-rpc/oms-rpc -v ./rpc/oms/oms.go
                    '''
                     
                }
    
            }
            stage('pms-rpc') {
                steps {
                    sh '''
                    mkdir -p target/pms-rpc
                    cp rpc/pms/etc/pms.yaml target/pms-rpc/pms-rpc.yaml
                    /usr/local/go/bin/go build -o target/pms-rpc/pms-rpc -v ./rpc/pms/pms.go
                    '''
                     
                }
    
            }
            stage('cms-rpc') {
                steps {
                    sh '''
                    mkdir -p target/cms-rpc
                    cp rpc/cms/etc/cms.yaml target/cms-rpc/cms-rpc.yaml
                    /usr/local/go/bin/go build -o target/cms-rpc/cms-rpc -v ./rpc/cms/cms.go
                    '''
                     
                }
    
            }
            stage('sms-rpc') {
                steps {
                    sh '''
                    mkdir -p target/sms-rpc
                    cp rpc/sms/etc/sms.yaml target/sms-rpc/sms-rpc.yaml
                    /usr/local/go/bin/go build -o target/sms-rpc/sms-rpc -v ./rpc/sms/sms.go
                    '''
                     
                }
    
            }
          
          }   
        }
        stage('build api') {
          parallel {
            stage('admin-api') {
                steps {
                    sh '''
                    mkdir -p target/admin-api
                    cp api/admin/etc/admin-api.yaml target/admin-api/admin-api.yaml
                    /usr/local/go/bin/go build -o target/admin-api/admin-api -v ./api/admin/admin.go
                    '''
                     
                }
    
            }
            stage('web-api') {
                steps {
                    sh '''
                    mkdir -p target/web-api
                    cp api/web/etc/web-api.yaml target/web-api/web-api.yaml
                    /usr/local/go/bin/go build -o target/web-api/web-api -v ./api/web/web.go
                    '''
                     
                }
    
            }
           stage('front-api') {
                steps {
                    sh '''
                    mkdir -p target/front-api
                    cp api/front/etc/front-api.yaml target/front-api/front-api.yaml
                    /usr/local/go/bin/go build -o target/front-api/front-api -v ./api/front/front.go
                    '''
                     
                }
    
            }
          }   
        }
       stage('stop') {
            steps {
                echo 'begin stop service'
            }
        }
        stage('stop service') {

          parallel {
            stage('stop sys-rpc') {
                steps {
                    sh '''
                    	pkill -f sys-rpc|| true
                    '''
                }
            }
            stage('stop ums-rpc') {
                steps {
                    sh '''
                    	pkill -f sys-rpc|| true
                    '''
                     
                }
    
            }
            stage('stop oms-rpc') {
                steps {
                    sh '''
                    	pkill -f sys-rpc|| true
                    '''
                     
                }
    
            }
            stage('stop pms-rpc') {
                steps {
                    sh '''
                    	pkill -f sys-rpc|| true
                    '''
                     
                }
    
            }
            stage('stop cms-rpc') {
                steps {
                    sh '''
                    	pkill -f sys-rpc|| true
                    '''
                     
                }
    
            }
            stage('stop sms-rpc') {
                steps {
                    sh '''
                    	pkill -f sys-rpc|| true
                    '''
                     
                }
    
            }
          
          }
        }
        stage('stop api') {

          parallel {
            stage('stop admin-api') {
                steps {
                    sh '''
                    	pkill -f sys-rpc|| true
                    '''
                }
    
            }
            stage('stop web-api') {
                steps {
                    sh '''
                    	pkill -f sys-rpc|| true
                    '''
                }
    
            }
           stage('stop front-api') {
                steps {
                    sh '''
                    	pkill -f sys-rpc|| true
                    '''
                     
                }
    
            }
          }
        }
       stage('start') {
            steps {
                echo 'begin start service'
            }
        }
        stage('start service') {
          parallel {
            stage('start sys-rpc') {
                steps {
                  withEnv(['JENKINS_NODE_COOKIE=dontkillme]']) {
                    sh '''
                    nohup ./target/sys-rpc/sys-rpc -f ./target/sys-rpc/sys-rpc.yaml  > /dev/null 2>&1 &
                    '''
                }
              }
            }
            stage('start ums-rpc') {
                steps {
                  withEnv(['JENKINS_NODE_COOKIE=dontkillme]']) {
                    sh '''
                    nohup ./target/ums-rpc/ums-rpc -f ./target/ums-rpc/ums-rpc.yaml  > /dev/null 2>&1 &
                    '''
                }
              }
    
            }
            stage('start oms-rpc') {
                steps {
                  withEnv(['JENKINS_NODE_COOKIE=dontkillme]']) {
                    sh '''
                    nohup ./target/oms-rpc/oms-rpc -f ./target/oms-rpc/oms-rpc.yaml  > /dev/null 2>&1 &
                    '''
                }
              }
    
            }
            stage('start pms-rpc') {
                steps {
                  withEnv(['JENKINS_NODE_COOKIE=dontkillme]']) {
                    sh '''
                    nohup ./target/pms-rpc/pms-rpc -f ./target/pms-rpc/pms-rpc.yaml  > /dev/null 2>&1 &
                    '''
                }
              }
    
            }
            stage('start cms-rpc') {
                steps {
                  withEnv(['JENKINS_NODE_COOKIE=dontkillme]']) {
                    sh '''
                    nohup ./target/cms-rpc/cms-rpc -f ./target/cms-rpc/cms-rpc.yaml  > /dev/null 2>&1 &
                    '''
                }
              }
    
            }
            stage('start sms-rpc') {
                steps {
                  withEnv(['JENKINS_NODE_COOKIE=dontkillme]']) {
                    sh '''
                    nohup ./target/sms-rpc/sms-rpc -f ./target/sms-rpc/sms-rpc.yaml  > /dev/null 2>&1 &
                    '''
                }
              }
    
            }
          }
        }
        stage('start api') {
          parallel {
            
            stage('start admin-api') {
                steps {
                  withEnv(['JENKINS_NODE_COOKIE=dontkillme]']) {
                    sh '''
                    nohup ./target/admin-api/admin-api -f ./target/admin-api/admin-api.yaml > /dev/null 2>&1 &
                    '''
                }
              }
    
            }
            stage('start web-api') {
                steps {
                  withEnv(['JENKINS_NODE_COOKIE=dontkillme]']) {
                    sh '''
                    nohup ./target/front-api/front-api -f ./target/front-api/front-api.yaml  > /dev/null 2>&1 &
                    '''
                }
              }
    
            }
           stage('start front-api') {
                steps {
                  withEnv(['JENKINS_NODE_COOKIE=dontkillme]']) {
                    sh '''
                    nohup ./target/web-api/web-api -f ./target/web-api/web-api.yaml  > /dev/null 2>&1 &
                    '''
                }
              }
    
            }
          }
        }
    }
}
```