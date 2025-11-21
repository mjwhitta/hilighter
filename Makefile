-include gomk/main.mk
-include local/Makefile

ifneq ($(unameS),windows)
spellcheck:
	@codespell -f -L hilight,hilights -S "*.pem,.git,go.*,gomk"
endif
