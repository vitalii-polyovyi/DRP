db.createUser(
    {
        user: process.env.MONGO_INITDB_USER,
        pwd: process.env.MONGO_INITDB_USER_PASSWORD,
        roles: [
            {
                role: "readWrite",
                db: process.env.MONGO_INITDB_DATABASE
            }
        ]
    }
);

db.apps.createIndex(
    {
        "app": 1
    },
    {
        unique: true,
        sparse: true,
    }
)
