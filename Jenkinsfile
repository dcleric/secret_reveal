node {
    timeout(time: 30, unit: 'MINUTES', activity: true) {
        // Tool versions

//        stage ('Checkout branch') {
//           // Check if build trigger is a PR and set branch name for checkout
//            def scmVars = checkout([
//                $class: 'GitSCM',
//                branches: [[name: branchName]],
//                doGenerateSubmoduleConfigurations: false,
//                extensions: [[$class: 'CleanBeforeCheckout']],
 //               userRemoteConfigs: [[ url: 'https://github.com/dcleric/secret_reveal.git']]
//            ])
//        }

// docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.43.0 golangci-lint run -v
        stage ('Go lint') {
            checkout scm
            sh(script: "docker run --rm -v ${env.WORKSPACE}:/app -w /app golangci/golangci-lint:v1.43.0 golangci-lint run -v")

//            docker.image('golangci/golangci-lint:v1.43.0').withRun('--rm -v $(pwd):/app -w /app').inside('golangci-lint run -v') { c ->
//               sh 'golangci-lint run -v'}
            }
    }
}
