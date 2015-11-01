#!/bin/bash

# exit on error
set -e
# setup bash
shopt -s extglob

# NOTE: currently will have problems if target contains a $1 directory - amongst other things!

GITUSER=aliwatters

# Targets DIR is working space, repo is where the code will be moved to.
TARGETDIR=deprecate
TARGETREPO=archive

REPONAME=$1
REPO="git@github.com:$GITUSER/$1.git"

echo Adding repo $1 to $TARGETREPO pulling from $REPO

mkdir -p $TARGETDIR
# clean old
rm -rf $TARGETDIR/$REPONAME

cd ~/$TARGETDIR

CMD="git clone $REPO"
echo "\$ $CMD"
$CMD

cd $REPONAME

# remove origin to ensure we don't check broken stuff in.
git remote rm origin

# Step 3
mkdir -p $REPONAME

# for all the files * and those .* (but not .. .)
for f in * .[^.]*; do
	echo "Checking $f"
	if [ "$f" != ".git" -a "$f" != "$REPONAME" ]; then
		CMD="git mv $f $REPONAME/."
		echo "\$ $CMD"
		$CMD
	fi
done

# review the changes.
git status

# commit the changes
git commit -m "Prepare $REPONAME to merge into $TARGETREPO"
cd ..

# clone our legacy destination if we don't have it.
if [ ! -d $TARGETREPO ]; then
	CMD="git clone git@github.com:$GITUSER/$TARGETREPO.git"
	echo "\$ $CMD"
	$CMD
fi

cd $TARGETREPO

pwd
CMD="git remote add $REPONAME ../$REPONAME"
echo "\$ $CMD"
$CMD

# don't promte for a message - just pull in.

GIT_MERGE_AUTOEDIT=no

CMD="git pull $REPONAME master"
echo "\$ $CMD"
$CMD

CMD="git remote rm $REPONAME"
echo "\$ $CMD"
$CMD

echo "We have added another repo to $TARGETREPO  - when you are done commit and push $TARGETREPO"
