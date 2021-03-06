#!/bin/bash

SCRIPT_FOLDER="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
echo $SCRIPT_FOLDER
TEMPDIR=`mktemp -d`

if [[ -z "$BRANCH" ]]; then
  BRANCH="feature/string-type-go";
fi

git clone -b $BRANCH https://github.com/VirgilSecurity/virgil-crypto-c.git $TEMPDIR

RETRES=$?
echo $RETRES
if [ "$RETRES" == "0" ]; then
  rm  -rf $SCRIPT_FOLDER/../{foundation,phe,sdk};
  cp -R $TEMPDIR/wrappers/go/{foundation,phe,sdk} $SCRIPT_FOLDER/../;
  for i in $(grep -R "virgil/foundation" $SCRIPT_FOLDER/../{foundation,phe,sdk} | cut -d ":" -f 1)
  do
  	echo  $i
    sed -i -e 's/virgil\/foundation/github.com\/VirgilSecurity\/virgil-commkit-go\/crypto\/wrapper\/foundation/g' $i
  done;
  gofmt -l -s -w $SCRIPT_FOLDER/../
fi
rm -rf $TEMPDIR