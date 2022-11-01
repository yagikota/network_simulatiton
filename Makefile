# change lamda by 0.1 in the range of 0.7~1.0 and simulate MM1K queue
.PHONY: range-simulate-MM1K
range-simulate-MM1K: ## range-simulate
	$(eval num:=$(shell seq 0.7 0.005 1.0 ))
	@for i in $(num); do \
		go run main.go simulate -l $$i -k 50; \
	done

# change lamda by 0.1 in the range of 0.7~1.0 and simulate MD1K queue
.PHONY: range-simulate-MD1K
range-simulate-MD1K: ## range-simulate
	$(eval num:=$(shell seq 0.7 0.005 1.0 ))
	@for i in $(num); do \
		go run main.go simulate -l $$i -k 50 -q 1; \
	done

# eval echo $$i
