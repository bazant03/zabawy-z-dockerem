const express = require("express")

const app = express()
const port = 8080

app.get("/", (req, res) => {
    res.send(
        '<h1 style="color:orange;text-align:center">Poważna aplikacja webowa jest modyfikowana<\h1>',
    )
})

app.listen(port, () => {
    console.log(`Aplikacja nasłuchuje na porcie ${port}`)
})
