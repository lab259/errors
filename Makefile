COVERDIR=$(CURDIR)/.cover
COVERAGEFILE=$(COVERDIR)/cover.out

.PHONY: test test-watch coverage coverage-html coverage-ci vet fmt

test:
	@ginkgo --failFast ./...

test-watch:
	@ginkgo watch -cover -r ./...

coverage-ci:
	@mkdir -p $(COVERDIR)
	@ginkgo -r -covermode=count --cover --trace ./
	@echo "mode: count" > "${COVERAGEFILE}"
	@find . -type f -name '*.coverprofile' -exec cat {} \; -exec rm -f {} \; | grep -h -v "^mode:" >> ${COVERAGEFILE}

coverage: coverage-ci
	@sed -i -e "s|_$(CURDIR)/|./|g" "${COVERAGEFILE}"
	@cp "${COVERAGEFILE}" coverage.txt

coverage-html:
	@go tool cover -html="${COVERAGEFILE}" -o .cover/report.html

vet:
	@go vet ./...

fmt:
	@go fmt ./...
