const express = require('express')
const axios = require('axios')
const cors = require('cors')
const redis = require('redis')
const DEFAULT_EXPIRY = 3600
const app = express()
app.use(cors())

const redisClient = redis.createClient()

app.get("/photos", async (req, res) => {
  const albumID = req.query.albumID
  redisClient.get('photos', async (error, photos) => {
    if (error) {
      throw new Error(error)
    }
    if (photos !== null) {
      console.log("cache HIT!")
      return res.json(JSON.parse(photos))
    }
    console.log("Cached has been missed. but we got it now")
    // else 
    const { data } = await axios.get(
      `https://jsonplaceholder.typicode.com/photos/`,
      { params: { albumID } }
    )

    redisClient.setEx('photos', DEFAULT_EXPIRY, JSON.stringify(data))
  })

  return res.json(data)
})

app.get("/photos/:id", async (req, res) => {
  const { data } = await axios.get(
    `https://jsonplaceholder.typicode.com/photos/${req.params.id}`
  )
  return res.json(data)
})
console.log("listening on port 3000")
app.listen(3000)
