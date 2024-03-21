FROM node:21-bullseye-slim

# set working directory
WORKDIR /home/app

# expose port
EXPOSE 5173

# stop signal
STOPSIGNAL SIGKILL

# healthcheck
HEALTHCHECK CMD curl -f http://localhost:5173/ || exit 1

# copy package.json and package-lock.json
COPY package*.json ./

# install dependencies
RUN npm install

# copy source code
COPY . .

# start app
CMD npm run dev
