# Use a slim version of Python for smaller image size
FROM python:3.12-slim

# Set working directory
WORKDIR /usr/src/app

# Install dependencies first for caching
COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

# Copy application code
COPY . .

# Use a non-root user for security
RUN useradd -m appuser
USER appuser

# Expose the port your app runs on
EXPOSE 5003

# Run the application
CMD ["python", "app.py"]
