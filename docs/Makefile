OUTPUT_DIR := ./bin
OUTPUT_NAME := go-resume-design.pdf
OUTPUT := $(OUTPUT_DIR)/$(OUTPUT_NAME)

OPT_SRC := $(shell find -type "f" -name "*.tex")

INPUT_SRC := ./design.md

all: $(OUTPUT)

$(OUTPUT): $(OUTPUT_DIR) $(INPUT_SRC) $(OPT_SRC)
	pandoc $(INPUT_SRC) \
		--pdf-engine=xelatex \
		-H $(OPT_SRC) \
		-V linkcolor:NavyBlue \
  	-V geometry:a4paper \
		-V geometry:"top=2cm, bottom=1.5cm, left=2cm, right=2cm" \
		--highlight-style zenburn \
		--toc \
		-N \
		-o $@


$(OUTPUT_DIR):
	mkdir -p $(OUTPUT_DIR)

.PHONY: clean
clean:
	rm $(OUTPUT)
