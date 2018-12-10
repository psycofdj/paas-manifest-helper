SOURCES = $(wildcard *.el | grep -v autoloader)
VERSION = $(shell cat paas-manifest-helper-pkg.el | head -n1 | awk '{print $$3}' | sed 's/"//g')

all: install

dist: emacs-paas-manifest-helper-$(VERSION).tar

emacs-paas-manifest-helper-$(VERSION).tar: $(SOURCES)
	@mkdir paas-manifest-helper-$(VERSION)
	@cp --parents -dR $(SOURCES) paas-manifest-helper-$(VERSION)/
	@tar cvf emacs-paas-manifest-helper-$(VERSION).tar paas-manifest-helper-$(VERSION)
	@rm -rf paas-manifest-helper-$(VERSION)

install: emacs-paas-manifest-helper-$(VERSION).tar
	@rm -rf ~/.emacs.d/elpa/paas-manifest-helper*/
	emacs --batch --eval "(defconst pkg-to-install \"$(PWD)/emacs-paas-manifest-helper-$(VERSION).tar\")" -l vendor/emacs-pkg-install.el
