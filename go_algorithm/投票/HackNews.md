- 每个帖子前面有一个向上的三角形，如果你觉得这个内容很好，就点击一下，投上一票。根据得票数，系统自动统计出热门文章排行榜。但是，并非得票最多的文章排在第一位，还要考虑时间因素，新文章应该比旧文章更容易得到好的排名。


### 实现
1. 公式 `Score = (P-1) / (T+2)^G`，
   - `P` 为得票数
     - 减1是为了忽略发帖人的投票，避免新文章得票数为0时，`Score` 为无穷大。
   - `T` 为文章发布时间距离现在的时间差
     - ，加上2是为了防止最新的帖子导致分母过小（之所以选择2，可能是因为从原始文章出现在其他网站，到转贴至Hacker News，平均需要两个小时）。
   - `G` 为一个常数("重力因子")，即帖子往下拉的力量，默认1.8，用来调整时间因素的重要性。
     - `G` 的值越大，时间因素的重要性越大，也就是说，新文章比旧文章更容易得到好的排名。`G` 的值越小，时间因素的重要性越小，也就是说，新文章比旧文章更难得到好的排名。

