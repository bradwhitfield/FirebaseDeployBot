if [[ -z "${FIREBASE_DEPLOY_TOKEN}" ]]; then
  echo "You must specify the FIREBASE_DEPLOY_TOKEN environment variable to use this image."
  echo "See the README on Docker Hub for more info."
  exit 1
fi

if [[ -z "${GIT_REPO}" ]]; then
  echo "You mush specify the GIT_REPO environemtn variable to use this image."
  echo "See the README on Docker Hub for more info."
  exit 1
fi

echo "Cloning repo to build hugo site."
git clone $GIT_REPO hugo-site
if [[ ! -d hugo-site ]]; then
  echo "Failed to clone $GIT_REPO."
  exit 1
fi
cd hugo-site

echo "Building hugo site."
hugo

echo "Uploading to Firebase."
if [[ -z "${FIREBASE_PROJECT}" ]]; then
  firebase deploy --token $FIREBASE_DEPLOY_TOKEN --project $FIREBASE_PROJECT
else
  firebase deploy --token $FIREBASE_DEPLOY_TOKEN
fi
