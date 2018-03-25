package controllers

import (
  "github.com/kataras/iris"
  "strconv"
  "../services"
)

func IncrementPairingCount(ctx iris.Context) {
  first, _ := strconv.Atoi(ctx.PostValue("user1"))
  second, _ := strconv.Atoi(ctx.PostValue("user2"))

  first, second = services.Reorder(first, second)

  pairCount, err := services.CreatePair(first, second)

  if err == nil {
    if saveErr := pairCount.IncrementCount(); saveErr == nil {
      ctx.JSON(iris.Map{"status": "success", "count": pairCount.Count})
      return
    }
  }
  ctx.JSON(iris.Map{"status": "error"})
}

func AllPairs(ctx iris.Context) {
  pairs := services.GetAllPairCounts()
  ctx.JSON(pairs)
}
