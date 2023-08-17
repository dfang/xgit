# get next_tag
next_tag=$(git describe --tags --abbrev=0 | awk -F. '{OFS="."; $NF+=1; print $0}')
echo $next_tag > version.txt

new_version=$(cat version.txt)
gsed -i "s/version=\"v[0-9.]*\"/version=\"$new_version\"/" scripts/install.sh


git tag -a $next_tag -m "$next_tag"
git push origin $next_tag