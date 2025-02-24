# Step 1: Build the front-end application
FROM node:20.5.0 AS build

WORKDIR /app

# Copy package.json and yarn.lock (for Yarn)
COPY package*.json yarn.lock ./

# Install dependencies using Yarn
RUN yarn install

# Copy the rest of the application code
COPY . .

# Build the Vue.js application for production
RUN yarn build

# Step 2: Serve the application with Nginx (optional, but recommended for production)
FROM nginx:alpine

# Copy the built Vue.js app from the build stage to the Nginx default directory
COPY --from=build /app/dist /usr/share/nginx/html

# Expose port 80 (default Nginx port)
EXPOSE 80

# Run Nginx in the foreground
CMD ["nginx", "-g", "daemon off;"]
