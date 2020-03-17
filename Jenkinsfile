node('haimaxy-jnlp') {
  stage('Prepare Stage') {
    echo "1. Prepare"
    checkout scm
    script {
        build_tag = sh(returnStdout: true, script: 'git rev-parse --short HEAD').trim()
        //当使用多流水线pipeline时，jenkins会根据git多分支branch名称注入一个叫BRANCH_NAME的环境变量，可以根据这个环境变量做版本判断等功能
        //注意，是自动注入的
        echo "####################### ${build_tag} ######## ${env.BRANCH_NAME} #### ${env.branch} #### ${env} #####"
        if (env.BRANCH_NAME != 'master') {
          build_tag = "${env.BRANCH_NAME}-${build_tag}"
        }
    }
  }
  stage('Test') {
    echo "2.Test Stage"
  }
  stage('Build') {
    echo "3.Build Stage"
    sh "docker build -t whatlong7/jenkins-demo:${build_tag} ."
  }  
  stage('Push') {
        echo "4.Push Docker Image Stage"
        withCredentials([usernamePassword(credentialsId: 'dockerHub', passwordVariable: 'dockerHubPassword', usernameVariable: 'dockerHubUser')]) {
            sh "docker login -u ${dockerHubUser} -p ${dockerHubPassword}"
            sh "docker push whatlong7/jenkins-demo:${build_tag}"
        }
    }
  stage('Deploy') {
        echo "5. Deploy Stage"
        if (env.BRANCH_NAME == 'master') {
            input "确认要部署线上环境吗？"
        }
        sh "sed -i 's/<BUILD_TAG>/${build_tag}/' k8s.yaml"
        sh "sed -i 's/<BRANCH_NAME>/${env.BRANCH_NAME}/' k8s.yaml"
        sh "kubectl apply -f k8s.yaml --record"
    }
}
