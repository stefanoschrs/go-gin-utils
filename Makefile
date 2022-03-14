bump:
ifeq ("$(git status | grep "nothing to commit, working tree clean")", "")
	echo "Please commit first"
	exit 1
endif
	git tag -a $$(git describe --tags --abbrev=0 | awk -F. '{OFS="."; $NF+=1; print $0}') -m "New Version"
	git push --follow-tags
