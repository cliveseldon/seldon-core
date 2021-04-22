VERSION := $(shell cat ../../version.txt)

aws-push:
	helm chart save . 709825985650.dkr.ecr.us-east-1.amazonaws.com/seldon-technologies/seldon-core:${VERSION}
	helm chart push 709825985650.dkr.ecr.us-east-1.amazonaws.com/seldon-technologies/seldon-core:${VERSION}
