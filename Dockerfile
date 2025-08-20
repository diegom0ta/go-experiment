FROM postgres:16

# Set environment variables (customize as needed)
ENV POSTGRES_USER=myuser
ENV POSTGRES_PASSWORD=mypassword
ENV POSTGRES_DB=wallet
ENV POSTGRES_HOST=localhost

# Expose PostgreSQL port
EXPOSE 5432
