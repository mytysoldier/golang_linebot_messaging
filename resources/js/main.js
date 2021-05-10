$(function () {
  // メッセージ送信ボタン押下時の処理
  $("#send").on("click", function (e) {
    e.preventDefault();
    // メッセージ送信
    $.post("/send", $("#form_message").serialize())
      .done(function () {
        alert("メッセージ送信しました。");
      })
      .fail(function () {
        alert("メッセージ送信失敗しました。");
      });
  });

  // クリアボタン押下時の処理
  $("#clear").on("click", function (e) {
    e.preventDefault();
    $("#message").val("");
  });
});
