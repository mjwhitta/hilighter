-include gomk/main.mk
-include local/Makefile

ifneq ($(unameS),windows)
spellcheck:
	@codespell -f -L hilight,hilighter,hilights -S "*.pem,.git,go.*,gomk"
endif
