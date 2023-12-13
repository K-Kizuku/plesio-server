.PHONY: docker/build_for_artifact_registry docker/push_for_artifact_registry

DOCKER_TAG = ''
_DOCKER_TAG = ''

ifneq ($(DOCKER_TAG),'')
  _DOCKER_TAG = ':$(DOCKER_TAG)'
endif

docker/build_for_artifact_registry:
	docker build ./ -t us-central1-docker.pkg.dev/hallowed-spider-407610/test-repo/udp-server-test-py$(DOCKER_TAG)

docker/push_for_artifact_registry:
	docker push us-central1-docker.pkg.dev/hallowed-spider-407610/test-repo/udp-server-test-py$(DOCKER_TAG)

