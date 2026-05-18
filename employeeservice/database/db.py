import os
import logging
from pymongo import MongoClient

client = None
db = None


def connect_db():
    global client, db

    mongodb_uri = "mongodb://"  + os.getenv("DATABASE_USERNAME") + ":" + os.getenv("DATABASE_PASSWORD") + "@" + os.getenv("DATABASE_URL") +"?authSource=admin"
    # mongodb://root:rootp@mongodb:27017/kindergarten?authSource=admin
   # mongodb_uri = os.getenv("MONGODB_URI")
    database_name = os.getenv("DATABASE_NAME", "kindergarten")

    if not mongodb_uri:
        raise RuntimeError("MONGODB_URI is not set")

    client = MongoClient(mongodb_uri, serverSelectionTimeoutMS=10000)
    db = client[database_name]

    # Test connection
    client.admin.command("ping")
    logging.info("Connected to MongoDB successfully!")


def get_collection(name):
    return db[name]
