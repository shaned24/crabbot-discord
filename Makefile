# Tokens
DISCORD_TOKEN ?="abcdefg"

# Docker
IMAGE         ?=shaned24/crabbot-discord
TAG           ?=latest
COMMAND		  ?=/opt/bin/demo

# Helm
NAME          ?=demo
RELEASE_NAME  ?=crabbot$(NAME)
CHART 		  ?=./deploy/crabbot

.PHONY: deploy deploy-dry-run
docker-build:
	@docker build . -t $(IMAGE):$(TAG)

docker-publish: docker-build
	@docker push $(IMAGE):$(TAG)

deploy:
	RELEASE_NAME=$(RELEASE_NAME) \
	CHART=$(CHART) \
	IMAGE=$(IMAGE) \
	TAG=$(TAG) \
	COMMAND=$(COMMAND) \
	DISCORD_TOKEN=$(DISCORD_TOKEN) \
	./scripts/deploy.sh

#	@helm upgrade --install $(RELEASE_NAME) $(CHART) \
#	--set "image.repository=$(IMAGE)" \
#	--set "image.tag=$(TAG)" \
#	--set "image.command=$(COMMAND)" \
#	--set "image.args={-t,$(DISCORD_TOKEN)}"

deploy-dry-run:
	helm upgrade --install --dry-run --debug \
	$(RELEASE_NAME) $(CHART) \
	--set "image.repository=$(IMAGE)" \
	--set "image.tag=$(TAG)" \
	--set "image.command=$(COMMAND)" \
	--set "image.args={-t,$(DISCORD_TOKEN)}"
