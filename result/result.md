# 出力結果の日付はいい感じに変更する

概要
同じメール送信処理をSync / Goroutine / NATS(JetStream)で比較しました

前提（共通条件）

- 1件=200ms（Sleep）
- 20件
- Goroutine worker数 = N
- NATS worker数 = N
測定指標
- 処理時間
- 1秒あたりの処理件数

実装概要

- Sync
- goroutine
- NATS(JetStream)
- それぞれのコード等を貼る

出力

```
2025/12/26 17:47:40 同期: Sending to user1 (user1@example.com)...
2025/12/26 17:47:40 同期: Done: user1
2025/12/26 17:47:40 同期: Sending to user2 (user2@example.com)...
2025/12/26 17:47:41 同期: Done: user2
2025/12/26 17:47:41 同期: Sending to user3 (user3@example.com)...
2025/12/26 17:47:41 同期: Done: user3
2025/12/26 17:47:41 同期: Sending to user4 (user4@example.com)...
2025/12/26 17:47:41 同期: Done: user4
2025/12/26 17:47:41 同期: Sending to user5 (user5@example.com)...
2025/12/26 17:47:41 同期: Done: user5
2025/12/26 17:47:41 同期: Sending to user6 (user6@example.com)...
2025/12/26 17:47:41 同期: Done: user6
2025/12/26 17:47:41 同期: Sending to user7 (user7@example.com)...
2025/12/26 17:47:42 同期: Done: user7
2025/12/26 17:47:42 同期: Sending to user8 (user8@example.com)...
2025/12/26 17:47:42 同期: Done: user8
2025/12/26 17:47:42 同期: Sending to user9 (user9@example.com)...
2025/12/26 17:47:42 同期: Done: user9
2025/12/26 17:47:42 同期: Sending to user10 (user10@example.com)...
2025/12/26 17:47:42 同期: Done: user10
2025/12/26 17:47:42 同期: Sending to user11 (user11@example.com)...
2025/12/26 17:47:42 同期: Done: user11
2025/12/26 17:47:42 同期: Sending to user12 (user12@example.com)...
2025/12/26 17:47:43 同期: Done: user12
2025/12/26 17:47:43 同期: Sending to user13 (user13@example.com)...
2025/12/26 17:47:43 同期: Done: user13
2025/12/26 17:47:43 同期: Sending to user14 (user14@example.com)...
2025/12/26 17:47:43 同期: Done: user14
2025/12/26 17:47:43 同期: Sending to user15 (user15@example.com)...
2025/12/26 17:47:43 同期: Done: user15
2025/12/26 17:47:43 同期: Sending to user16 (user16@example.com)...
2025/12/26 17:47:43 同期: Done: user16
2025/12/26 17:47:43 同期: Sending to user17 (user17@example.com)...
2025/12/26 17:47:44 同期: Done: user17
2025/12/26 17:47:44 同期: Sending to user18 (user18@example.com)...
2025/12/26 17:47:44 同期: Done: user18
2025/12/26 17:47:44 同期: Sending to user19 (user19@example.com)...
2025/12/26 17:47:44 同期: Done: user19
2025/12/26 17:47:44 同期: Sending to user20 (user20@example.com)...
2025/12/26 17:47:44 同期: Done: user20

同期処理完了: 4.023804708s, rps: 4.970420149923439
```

```
2025/12/26 17:47:44 非同期: Sending to user1 (user1@example.com)...
2025/12/26 17:47:44 非同期: Sending to user2 (user2@example.com)...
2025/12/26 17:47:44 非同期: Done: user2
2025/12/26 17:47:44 非同期: Sending to user3 (user3@example.com)...
2025/12/26 17:47:44 非同期: Done: user1
2025/12/26 17:47:44 非同期: Sending to user4 (user4@example.com)...
2025/12/26 17:47:45 非同期: Done: user4
2025/12/26 17:47:45 非同期: Sending to user5 (user5@example.com)...
2025/12/26 17:47:45 非同期: Done: user3
2025/12/26 17:47:45 非同期: Sending to user6 (user6@example.com)...
2025/12/26 17:47:45 非同期: Done: user6
2025/12/26 17:47:45 非同期: Sending to user7 (user7@example.com)...
2025/12/26 17:47:45 非同期: Done: user5
2025/12/26 17:47:45 非同期: Sending to user8 (user8@example.com)...
2025/12/26 17:47:45 非同期: Done: user8
2025/12/26 17:47:45 非同期: Sending to user9 (user9@example.com)...
2025/12/26 17:47:45 非同期: Done: user7
2025/12/26 17:47:45 非同期: Sending to user10 (user10@example.com)...
2025/12/26 17:47:45 非同期: Done: user9
2025/12/26 17:47:45 非同期: Sending to user11 (user11@example.com)...
2025/12/26 17:47:45 非同期: Done: user10
2025/12/26 17:47:45 非同期: Sending to user12 (user12@example.com)...
2025/12/26 17:47:45 非同期: Done: user12
2025/12/26 17:47:45 非同期: Sending to user13 (user13@example.com)...
2025/12/26 17:47:45 非同期: Done: user11
2025/12/26 17:47:45 非同期: Sending to user14 (user14@example.com)...
2025/12/26 17:47:46 非同期: Done: user14
2025/12/26 17:47:46 非同期: Sending to user15 (user15@example.com)...
2025/12/26 17:47:46 非同期: Done: user13
2025/12/26 17:47:46 非同期: Sending to user16 (user16@example.com)...
2025/12/26 17:47:46 非同期: Done: user16
2025/12/26 17:47:46 非同期: Sending to user17 (user17@example.com)...
2025/12/26 17:47:46 非同期: Done: user15
2025/12/26 17:47:46 非同期: Sending to user18 (user18@example.com)...
2025/12/26 17:47:46 非同期: Done: user18
2025/12/26 17:47:46 非同期: Sending to user19 (user19@example.com)...
2025/12/26 17:47:46 非同期: Done: user17
2025/12/26 17:47:46 非同期: Sending to user20 (user20@example.com)...
2025/12/26 17:47:46 非同期: Done: user20
2025/12/26 17:47:46 非同期: Done: user19

非同期処理(Goroutine)完了: 2.01336525s, rps: 9.933617360287707
```

```
created the stream
2026/01/12 18:01:25 published stream=Tasks, seq=373
2026/01/12 18:01:25 published stream=Tasks, seq=374
2026/01/12 18:01:25 published stream=Tasks, seq=375
2026/01/12 18:01:25 published stream=Tasks, seq=376
2026/01/12 18:01:25 published stream=Tasks, seq=377
2026/01/12 18:01:25 published stream=Tasks, seq=378
2026/01/12 18:01:25 published stream=Tasks, seq=379
2026/01/12 18:01:25 published stream=Tasks, seq=380
2026/01/12 18:01:25 published stream=Tasks, seq=381
2026/01/12 18:01:25 published stream=Tasks, seq=382
2026/01/12 18:01:25 published stream=Tasks, seq=383
2026/01/12 18:01:25 published stream=Tasks, seq=384
2026/01/12 18:01:25 published stream=Tasks, seq=385
2026/01/12 18:01:25 published stream=Tasks, seq=386
2026/01/12 18:01:25 published stream=Tasks, seq=387
2026/01/12 18:01:25 published stream=Tasks, seq=388
2026/01/12 18:01:25 published stream=Tasks, seq=389
2026/01/12 18:01:25 published stream=Tasks, seq=390
2026/01/12 18:01:25 published stream=Tasks, seq=391
2026/01/12 18:01:25 published stream=Tasks, seq=392
2026/01/12 18:01:26 msg data: "{\"user_id\":\"user1\"}" on subject "tasks.results.1768208485986239000"
2026/01/12 18:01:26 msg data: "{\"user_id\":\"user2\"}" on subject "tasks.results.1768208485986239000"
2026/01/12 18:01:26 msg data: "{\"user_id\":\"user3\"}" on subject "tasks.results.1768208485986239000"
2026/01/12 18:01:26 msg data: "{\"user_id\":\"user4\"}" on subject "tasks.results.1768208485986239000"
2026/01/12 18:01:26 msg data: "{\"user_id\":\"user5\"}" on subject "tasks.results.1768208485986239000"
2026/01/12 18:01:26 msg data: "{\"user_id\":\"user6\"}" on subject "tasks.results.1768208485986239000"
2026/01/12 18:01:26 msg data: "{\"user_id\":\"user7\"}" on subject "tasks.results.1768208485986239000"
2026/01/12 18:01:26 msg data: "{\"user_id\":\"user8\"}" on subject "tasks.results.1768208485986239000"
2026/01/12 18:01:26 msg data: "{\"user_id\":\"user10\"}" on subject "tasks.results.1768208485986239000"
2026/01/12 18:01:26 msg data: "{\"user_id\":\"user9\"}" on subject "tasks.results.1768208485986239000"
2026/01/12 18:01:27 msg data: "{\"user_id\":\"user11\"}" on subject "tasks.results.1768208485986239000"
2026/01/12 18:01:27 msg data: "{\"user_id\":\"user12\"}" on subject "tasks.results.1768208485986239000"
2026/01/12 18:01:27 msg data: "{\"user_id\":\"user13\"}" on subject "tasks.results.1768208485986239000"
2026/01/12 18:01:27 msg data: "{\"user_id\":\"user14\"}" on subject "tasks.results.1768208485986239000"
2026/01/12 18:01:27 msg data: "{\"user_id\":\"user15\"}" on subject "tasks.results.1768208485986239000"
2026/01/12 18:01:27 msg data: "{\"user_id\":\"user16\"}" on subject "tasks.results.1768208485986239000"
2026/01/12 18:01:27 msg data: "{\"user_id\":\"user17\"}" on subject "tasks.results.1768208485986239000"
2026/01/12 18:01:27 msg data: "{\"user_id\":\"user18\"}" on subject "tasks.results.1768208485986239000"
2026/01/12 18:01:28 msg data: "{\"user_id\":\"user19\"}" on subject "tasks.results.1768208485986239000"
2026/01/12 18:01:28 msg data: "{\"user_id\":\"user20\"}" on subject "tasks.results.1768208485986239000"
total = 2.017793292s rps = 9.911818063472877
```

worker1

```
2026/01/12 18:01:18 waiting for messages...
2026/01/12 18:01:25 NATS: Sending to user1 (user1@example.com)...
2026/01/12 18:01:26 NATS: Sending to user3 (user3@example.com)...
2026/01/12 18:01:26 NATS: Sending to user5 (user5@example.com)...
2026/01/12 18:01:26 NATS: Sending to user7 (user7@example.com)...
2026/01/12 18:01:26 NATS: Sending to user9 (user9@example.com)...
2026/01/12 18:01:26 NATS: Sending to user11 (user11@example.com)...
2026/01/12 18:01:27 NATS: Sending to user13 (user13@example.com)...
2026/01/12 18:01:27 NATS: Sending to user15 (user15@example.com)...
2026/01/12 18:01:27 NATS: Sending to user17 (user17@example.com)...
2026/01/12 18:01:27 NATS: Sending to user19 (user19@example.com)...
```

worker2

```
2026/01/12 18:01:20 waiting for messages...
2026/01/12 18:01:25 NATS: Sending to user2 (user2@example.com)...
2026/01/12 18:01:26 NATS: Sending to user4 (user4@example.com)...
2026/01/12 18:01:26 NATS: Sending to user6 (user6@example.com)...
2026/01/12 18:01:26 NATS: Sending to user8 (user8@example.com)...
2026/01/12 18:01:26 NATS: Sending to user10 (user10@example.com)...
2026/01/12 18:01:26 NATS: Sending to user12 (user12@example.com)...
2026/01/12 18:01:27 NATS: Sending to user14 (user14@example.com)...
2026/01/12 18:01:27 NATS: Sending to user16 (user16@example.com)...
2026/01/12 18:01:27 NATS: Sending to user18 (user18@example.com)...
2026/01/12 18:01:27 NATS: Sending to user20 (user20@example.com)...
```

結果（表）

| mode | 件数 | メール送信処理時間（sleep） | 全体の処理時間 | 1秒あたりの処理件数 | worker数 |
| --- | --- | --- | --- | --- | --- |
| sync | 20 | 200ms | 4.02s | 4.97件 | 1 |
| goroutine | 20 | 200ms | 2.01s | 9.93件 | 2 |
| nats-jetstream | 20 | 200ms | 2.02s | 9.91件 | 2 |
| goroutine | 100 | 200ms | 2.015s | 49.6件 | 10 |

NATS JetStreamでワーカーを1つ止めた時の挙動

ワーカー1(途中で落とす)

```
2026/01/12 20:12:30 waiting for messages...
2026/01/12 20:12:48 NATS: Sending to user2 (user2@example.com)...
2026/01/12 20:12:48 NATS: Sending to user4 (user4@example.com)...
2026/01/12 20:12:48 NATS: Sending to user6 (user6@example.com)...
2026/01/12 20:12:48 NATS: Sending to user8 (user8@example.com)...
2026/01/12 20:12:48 NATS: Sending to user10 (user10@example.com)...
^C2026/01/12 20:12:49 Shutting down...
```

ワーカー2(user10から引き継ぎ)

```
2026/01/12 20:12:28 waiting for messages...
2026/01/12 20:12:48 NATS: Sending to user1 (user1@example.com)...
2026/01/12 20:12:48 NATS: Sending to user3 (user3@example.com)...
2026/01/12 20:12:48 NATS: Sending to user5 (user5@example.com)...
2026/01/12 20:12:48 NATS: Sending to user7 (user7@example.com)...
2026/01/12 20:12:48 NATS: Sending to user9 (user9@example.com)...
2026/01/12 20:12:49 NATS: Sending to user11 (user11@example.com)...
2026/01/12 20:12:49 NATS: Sending to user13 (user13@example.com)...
2026/01/12 20:12:49 NATS: Sending to user15 (user15@example.com)...
2026/01/12 20:12:49 NATS: Sending to user17 (user17@example.com)...
2026/01/12 20:12:49 NATS: Sending to user19 (user19@example.com)...
2026/01/12 20:13:18 NATS: Sending to user10 (user10@example.com)... // ここから処理引き継ぎ
2026/01/12 20:13:18 NATS: Sending to user12 (user12@example.com)...
2026/01/12 20:13:18 NATS: Sending to user14 (user14@example.com)...
2026/01/12 20:13:18 NATS: Sending to user16 (user16@example.com)...
2026/01/12 20:13:18 NATS: Sending to user18 (user18@example.com)...
2026/01/12 20:13:19 NATS: Sending to user20 (user20@example.com)...
```

producer側

```
created the stream
2026/01/12 20:12:48 published stream=Tasks, seq=493
2026/01/12 20:12:48 published stream=Tasks, seq=494
2026/01/12 20:12:48 published stream=Tasks, seq=495
2026/01/12 20:12:48 published stream=Tasks, seq=496
2026/01/12 20:12:48 published stream=Tasks, seq=497
2026/01/12 20:12:48 published stream=Tasks, seq=498
2026/01/12 20:12:48 published stream=Tasks, seq=499
2026/01/12 20:12:48 published stream=Tasks, seq=500
2026/01/12 20:12:48 published stream=Tasks, seq=501
2026/01/12 20:12:48 published stream=Tasks, seq=502
2026/01/12 20:12:48 published stream=Tasks, seq=503
2026/01/12 20:12:48 published stream=Tasks, seq=504
2026/01/12 20:12:48 published stream=Tasks, seq=505
2026/01/12 20:12:48 published stream=Tasks, seq=506
2026/01/12 20:12:48 published stream=Tasks, seq=507
2026/01/12 20:12:48 published stream=Tasks, seq=508
2026/01/12 20:12:48 published stream=Tasks, seq=509
2026/01/12 20:12:48 published stream=Tasks, seq=510
2026/01/12 20:12:48 published stream=Tasks, seq=511
2026/01/12 20:12:48 published stream=Tasks, seq=512
2026/01/12 20:12:48 msg data: "{\"user_id\":\"user1\"}" on subject "tasks.results.1768216368158676000"
2026/01/12 20:12:48 msg data: "{\"user_id\":\"user2\"}" on subject "tasks.results.1768216368158676000"
2026/01/12 20:12:48 msg data: "{\"user_id\":\"user4\"}" on subject "tasks.results.1768216368158676000"
2026/01/12 20:12:48 msg data: "{\"user_id\":\"user3\"}" on subject "tasks.results.1768216368158676000"
2026/01/12 20:12:48 msg data: "{\"user_id\":\"user6\"}" on subject "tasks.results.1768216368158676000"
2026/01/12 20:12:48 msg data: "{\"user_id\":\"user5\"}" on subject "tasks.results.1768216368158676000"
2026/01/12 20:12:48 msg data: "{\"user_id\":\"user8\"}" on subject "tasks.results.1768216368158676000"
2026/01/12 20:12:48 msg data: "{\"user_id\":\"user7\"}" on subject "tasks.results.1768216368158676000"
2026/01/12 20:12:49 msg data: "{\"user_id\":\"user9\"}" on subject "tasks.results.1768216368158676000"
2026/01/12 20:12:49 msg data: "{\"user_id\":\"user11\"}" on subject "tasks.results.1768216368158676000"
2026/01/12 20:12:49 msg data: "{\"user_id\":\"user13\"}" on subject "tasks.results.1768216368158676000"
2026/01/12 20:12:49 msg data: "{\"user_id\":\"user15\"}" on subject "tasks.results.1768216368158676000"
2026/01/12 20:12:49 msg data: "{\"user_id\":\"user17\"}" on subject "tasks.results.1768216368158676000"
2026/01/12 20:12:50 msg data: "{\"user_id\":\"user19\"}" on subject "tasks.results.1768216368158676000"
2026/01/12 20:13:18 msg data: "{\"user_id\":\"user10\"}" on subject "tasks.results.1768216368158676000"
2026/01/12 20:13:18 msg data: "{\"user_id\":\"user12\"}" on subject "tasks.results.1768216368158676000"
2026/01/12 20:13:18 msg data: "{\"user_id\":\"user14\"}" on subject "tasks.results.1768216368158676000"
2026/01/12 20:13:18 msg data: "{\"user_id\":\"user16\"}" on subject "tasks.results.1768216368158676000"
2026/01/12 20:13:19 msg data: "{\"user_id\":\"user18\"}" on subject "tasks.results.1768216368158676000"
2026/01/12 20:13:19 msg data: "{\"user_id\":\"user20\"}" on subject "tasks.results.1768216368158676000"
total = 31.227643459s rps = 0.6404581897528959
```

まとめ

- 複数ワーカーでは処理を同時に実行できるため、200msの待ちが重なり、全体の処理時間が短くなる
- 同時実行のため、実行毎に処理順が異なっていた。
- NATS JetStreamを使うことで、一つのワーカー（プロセス）が落ちても未Ackのジョブが残り、再配信されて別のワーカーが処理を引き継ぐことができる。

感想

- ワーカー数を増やすと全体の処理時間は短くなるが、CPUやリソースが詰まる可能性がある。特にメール送信のような外部のサービス・APIを使う際は制限に引っかかる可能性が高くなりそうだと思いました。ワーカー数はサービス全体のリソースやUX、外部との連携など考慮して適切に設定する必要があると思いました。
- また、非同期処理=速いだけでなく、処理を待たずに進める、分離して運用できるなど複数の価値があると思いました。
- NATS JetStreamはプロセスが落ちても別ワーカーで処理を再開できるため、goroutineよりも障害に強く、運用しやすいツールだと思いました。
- 今回はメール送信処理に失敗なしで比較したが、JetStreamには未完了時の再配信もできるため、次はメール送信に失敗を入れて、再配信の挙動も確認してみたいです