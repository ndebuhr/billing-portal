node {
   stage('Preparation') {
      deleteDir()
      git '${REPO}'
      sh "git submodule update --init --recursive"
   }
   stage('Build') {
      withEnv(["DOCKER_HOST=${DOCKERHOST}"]) {
        dir("payments/payment-cdi-event") {
          sh "mvn clean install"
        }
        dir("haproxy") {
          sh "docker build --no-cache -t ${REGISTRY}/billing-portal/haproxy:1.0.$BUILD_NUMBER ."
        }
        dir("billing-service") {
          sh "docker build --no-cache -t ${REGISTRY}/billing-portal/static-site:1.0.$BUILD_NUMBER ."
        }
      }
   }
   stage('Zip DB Scripts') {
      zip zipFile: 'sql/sql-scripts.zip', archive: false, dir: 'sql'
   }
   stage('Scan Artifacts') {
      withEnv(["DOCKER_HOST=${DOCKERHOST}"]) {
        sh 'docker run --rm -t -v $PWD/billing-service:/usr/src newtmitch/sonar-scanner:4.0 -Dsonar.projectKey=billing-service -Dsonar.host.url=${SONARHOST} -Dsonar.login=${SONARTOKEN}'
      }
   }
   stage('Push') {
      withEnv(["DOCKER_HOST=${DOCKERHOST}"]) {
        sh 'echo "${DOCKERPASSWORD}" | docker login -u ${DOCKERUSER} --password-stdin https://${REGISTRY}'
        sh 'docker push ${REGISTRY}/billing-portal/haproxy:1.0.$BUILD_NUMBER'
        sh 'docker push ${REGISTRY}/billing-portal/static-site:1.0.$BUILD_NUMBER'
      }
   }
   stage('Setup Deployment Package') {
       sh 'chmod +x ./xlw'
       sh './xlw apply -v --values BUILD_NUMBER=$BUILD_NUMBER,REGISTRY=${REGISTRY} --xl-deploy-url ${XLDHOST} --xl-deploy-username ${XLDUSER} --xl-deploy-password ${XLDPASSWORD} -f xl-deploy-billing-service.yaml'
   }
   stage('Archive Artifacts') {
      archiveArtifacts artifacts: 'payments/payment-cdi-event/target/payment-cdi-event.war, mongodb/scripts.zip', fingerprint: true
   }
   stage('Docker Cleanup') {
      withEnv(["DOCKER_HOST=${DOCKERHOST}"]) {
        sh 'docker rmi -f ${REGISTRY}/billing-portal/static-site:1.0.$BUILD_NUMBER'
        sh 'docker rmi -f ${REGISTRY}/billing-portal/haproxy:1.0.$BUILD_NUMBER'
      }
   }
}