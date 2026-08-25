const express = require('express')
const axios = require('axios')
const redis = require('redis')
const app = express();
const port = 3000;

// 1. Create and connect Redis client
const client = redis.createClient();
client.on("error", (err) => console.log("Redis error =>", err));

// 2. Define endpoint with caching logic
app.get("/api/photos", async (req, res) => {
  try {
    // 3. Attempt to get data from Redis cache
    const data = await client.get("photos");

    if (data) {
      console.log("Cache hit");
      res.status(200).json(JSON.parse(data));
    } else {
      console.log("Cache miss");
      // 4. Fetch from external API if cache miss
      const response = await axios.get("https://jsonplaceholder.typicode.com/photos");
      const photos = response.data;

      // 5. Store data in Redis with expiration (e.g., 10 seconds)
      await client.setEx("photos", 10, JSON.stringify(photos));
      res.status(200).json(photos);
    }
  } catch (error) {
    console.error(error);
    res.status(500).send("Internal Server Error");
  }
});
async function startserver() {

  await client.connect();
  app.listen(port, () => {
    console.log(`Server running at http://localhost:${port}`);
  });
}
startserver()

