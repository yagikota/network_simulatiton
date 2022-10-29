# change lamda by 0.1 in the range of 0.7~1.0 and simulate
.PHONY: range-simulate
range-simulate: ## range-simulate
	$(eval num:=$(shell seq 0.7 0.01 1.0 ))
	@for i in $(num); do \
		go run main.go simulate -l $$i -k 1000; \
	done

# eval echo $$i
