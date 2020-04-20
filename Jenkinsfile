node {
   stage('Preparation') {
      deleteDir()
      git '${REPO}'
      sh "git checkout ${BRANCH}"
      sh "git submodule update --init --recursive"
   }
   stage('Build') {
      withEnv(["DOCKER_HOST=${DOCKERHOST}"]) {
        dir("payments/payment-cdi-event") {
          sh "mvn clean install"
        }
        dir("haproxy") {
          sh "docker build --no-cache -t ${BUILDREGISTRY}/billing-portal/haproxy:1.0.$BUILD_NUMBER ."
        }
        dir("billing-service") {
          sh "docker build --no-cache -t ${BUILDREGISTRY}/billing-portal/static-site:1.0.$BUILD_NUMBER ."
        }
        dir("api") {
          sh "docker build --no-cache -t ${BUILDREGISTRY}/billing-portal/api:1.0.$BUILD_NUMBER ."
        }
        dir("scanner") {
          sh 'docker build --no-cache -t sonar-scanner:latest .'
        }
      }
   }
   stage('Zip DB Scripts') {
      zip zipFile: 'sql/sql-scripts.zip', archive: false, dir: 'sql'
   }
   stage('Scan Artifacts') {
      withEnv(["DOCKER_HOST=${DOCKERHOST}"]) {
        sh 'docker run --rm -t sonar-scanner -Dsonar.projectKey=billing-service -Dsonar.host.url=${SONARHOST} -Dsonar.login=${SONARTOKEN}'
      }
   }
   stage('Push') {
      withEnv(["DOCKER_HOST=${DOCKERHOST}"]) {
        sh 'echo "${DOCKERPASSWORD}" | docker login -u ${DOCKERUSER} --password-stdin https://${BUILDREGISTRY}'
        sh 'docker push ${BUILDREGISTRY}/billing-portal/haproxy:1.0.$BUILD_NUMBER'
        sh 'docker push ${BUILDREGISTRY}/billing-portal/static-site:1.0.$BUILD_NUMBER'
        sh 'docker push ${BUILDREGISTRY}/billing-portal/api:1.0.$BUILD_NUMBER'
      }
   }
   stage('Setup Deployment Package') {
       sh 'chmod +x ./xlw'
       sh './xlw apply -v --values BUILD_NUMBER=$BUILD_NUMBER,DEPLOYMENTREGISTRY=${DEPLOYMENTREGISTRY} --xl-deploy-url ${XLDHOST} --xl-deploy-username ${XLDUSER} --xl-deploy-password ${XLDPASSWORD} -f xl-deploy-billing-service.yaml'
   }
   stage('Archive Artifacts') {
      archiveArtifacts artifacts: 'payments/payment-cdi-event/target/payment-cdi-event.war, mongodb/scripts.zip', fingerprint: true
   }
   stage('Docker Cleanup') {
      withEnv(["DOCKER_HOST=${DOCKERHOST}"]) {
        sh 'docker rmi -f ${BUILDREGISTRY}/billing-portal/haproxy:1.0.$BUILD_NUMBER'
        sh 'docker rmi -f ${BUILDREGISTRY}/billing-portal/static-site:1.0.$BUILD_NUMBER'
        sh 'docker rmi -f ${BUILDREGISTRY}/billing-portal/api:1.0.$BUILD_NUMBER'
        sh 'docker rmi -f sonar-scanner:latest'
      }
   }
}