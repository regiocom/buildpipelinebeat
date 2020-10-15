# buildpipelinebeat

Welcome to buildpipelinebeat.

Ensure that this folder is at the following location:
`${GOPATH}/src/github.com/regiocom/buildpipelinebeat`

## Getting Started with buildpipelinebeat

### Requirements

- Docker
- Jenkins

For Development:

- Golang 1.14.7

### Usage

#### Local (for testing)

The Commandline to insert into the jenkinsfile (if you put the binary into a reachable path)

```sh
.\buildpipelinebeat -c buildpipelinebeat.yml -E "buildpipelinebeat.team=Teamname" -E "buildpipelinebeat.project=ProjectName" -E "buildpipelinebeat.pipeline=PipelineName" -E "buildpipelinebeat.status=Test" -d "*"
```

#### Docker

```bash
# substitute HUB and NAME accordingly
# Log yourself in before pushing
export DOCKER_HUB=localhost:5001/
export DOCKER_IMAGENAME=yourname/buildpipelinebeat:latest
git clone https://github.com/regiocom/buildpipelinebeat
cd buildpipelinebeat
docker build -t ${DOCKER_HUB}${DOCKER_IMAGENAME} .
docker push ${DOCKER_HUB}${DOCKER_IMAGENAME}
docker run --rm buildpipelinebeat -E "output.elasticsearch.hosts=[ 'localhost:5601' ]" -E "buildpipelinebeat.team=Teamname" -E "buildpipelinebeat.project=ProjectName" -E "buildpipelinebeat.pipeline=PipelineName" -E "buildpipelinebeat.status=Test" -d "*"
```

#### Jenkins

have a look in the examples directory

### Clone

To clone buildpipelinebeat from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/src/github.com/regiocom/buildpipelinebeat
git clone https://github.com/regiocom/buildpipelinebeat ${GOPATH}/src/github.com/regiocom/buildpipelinebeat
```

For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

### Build

To build the binary for buildpipelinebeat run the command below. This will generate a binary
in the same directory with the name buildpipelinebeat.

```
make
```

### Run

To run buildpipelinebeat with debugging output enabled, run:

```
./buildpipelinebeat -c buildpipelinebeat.yml -e -d "*"
```

### Test

This is only a basic test and no functionality \
To test buildpipelinebeat, run the following command:

```
make testsuite
```

alternatively:

```
make unit-tests
make system-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`

### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `fields.yml` by running the following command.

```
make update
```

### Cleanup

To clean buildpipelinebeat source code, run the following command:

```
make fmt
```

To clean up the build directory and generated artifacts, run:

```
make clean
```

## Packaging

The beat frameworks provides tools to crosscompile and package your beat for different platforms. This requires [docker](https://www.docker.com/) and vendoring as described above. To build packages of your beat, run the following command:

```
make release
```

This will fetch and create all images required for the build process. The whole process to finish can take several minutes.
