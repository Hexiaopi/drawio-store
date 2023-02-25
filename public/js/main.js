const drawModal = new bootstrap.Modal(document.getElementById('drawModal'), { keyboard: false })

function createImage () {
  var key = $("#inputDraw").val();
  if (key !== "") {
    $.ajax({
      type: "post",
      url: "/api/v1/image/" + key,
      async: true,
      success: function (data) {
        window.location.reload();
      }
    });
  }
}

function deleteImage (name) {
  $.ajax({
    type: "delete",
    url: "/api/v1/image/" + name,
    async: true,
    success: function (data) {
      window.location.reload();
    }
  });
}