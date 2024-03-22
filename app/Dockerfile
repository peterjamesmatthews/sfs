FROM node:21-bullseye-slim

# set working directory
WORKDIR /sfs/app

# expose port
EXPOSE 5173

# stop signal
STOPSIGNAL SIGKILL

# copy package.json and package-lock.json
COPY package*.json ./

# install dependencies
RUN npm install

# start app
CMD npm run dev
