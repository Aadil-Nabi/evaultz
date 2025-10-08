# ─── Stage 1: Builder ───
FROM node:20-alpine AS builder

WORKDIR /app

# Install dependencies separately (for better caching)
COPY package.json package-lock.json ./
RUN npm ci

# Copy source code
COPY . .

# Build Next.js app
RUN npm run build

# ─── Stage 2: Runtime ───
FROM node:20-alpine AS runner

WORKDIR /app

ENV NODE_ENV=production
# Next.js requires this for produ  ction
ENV NEXT_TELEMETRY_DISABLED=1

# Copy only the built files & necessary assets
COPY --from=builder /app/.next ./.next
COPY --from=builder /app/public ./public
COPY --from=builder /app/package.json ./package.json

# Install only production dependencies
RUN npm install --omit=dev --ignore-scripts --prefer-offline

# Expose Next.js default port
EXPOSE 3000

# Start Next.js
CMD ["npm", "run", "start"]
