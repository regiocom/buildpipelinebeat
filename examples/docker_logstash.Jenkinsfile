#!/usr/bin/env groovy

def buildpipelinebeat_DockerPath = 'localhost:5001/buildpipelinebeat:latest'
def buildpipelinebeat_TeamName = 'defaultTeam'
def buildpipelinebeat_ProjectName = 'defaultProject'
def buildpipelinebeat_PipelineName = 'defaultPipeline'
def buildpipelinebeat_LogstashHosts = 'hostip:port' /* For multientries use the following sheme:
  'hostip:port\", \"hostip2:port' */

// the parameter -d "*" (for the beat not docker!) activates the debug mode where the pushed message is printed out to the docker log
def buildpipelinebeat_BaseString = 'docker run --rm ${buildpipelinebeat_DockerPath} -E \"output.logstash.hosts=[ \"${buildpipelinebeat_LogstashHosts}\" ]\" -E \"buildpipelinebeat.team=${buildpipelinebeat_TeamName}\" -E \"buildpipelinebeat.project=${buildpipelinebeat_ProjectName}\" -E \"buildpipelinebeat.pipeline=${buildpipelinebeat_PipelineName}\" -E \"buildpipelinebeat.status='

def notifyStarted() {
  stage('Notify currently building') {
    container('docker') {
      echo "Running buildpipelinebeat"
      sh "${buildpipelinebeat_BaseString}Building\""
    }
  }
}

def notifySuccess() {
  stage('Notify successful build') {
    container('docker') {
      echo "Running buildpipelinebeat"
      sh "${buildpipelinebeat_BaseString}Success\""
    }
  }
}

def notifyFailure(error) {
  stage('Notify failed build') {
    container('docker') {
      echo "Running buildpipelinebeat"
      sh "${buildpipelinebeat_BaseString}Failure\" -E \"buildpipelinebeat.error=${error}\""
    }
  }
}

podTemplate(
  containers: [
    containerTemplate(name: 'docker', image: 'docker',  ttyEnabled: true, command: 'cat'
    ])
  ],
) {
    try {
      notifyStarted()

      /* ... existing build steps ... */

      notifySuccess()
    } catch (e) {
      notifyFailure(e)
      throw e
    }
}
