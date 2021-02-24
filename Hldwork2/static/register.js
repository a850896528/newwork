var uname = document.querySelector(".Username");
var pwd = document.querySelector(".password");
var tel = document.querySelector(".tlnumber");
var btn2 = document.querySelector(".btn2");
var flag1, flag2, flag3;
var utel = /^(13[0-9]|14[5|7]|15[0|1|2|3|5|6|7|8|9]|18[0|1|2|3|5|6|7|8|9])\d{8}$/;
var uPattern = /^[A-Za-z0-9_\u4e00-\u9fa5]{5,16}$/;
var upwd = /^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{6,20}$/;
var check = document.querySelector(".check");
uname.addEventListener("blur", function () {
  if (uPattern.test(this.value)) {
    this.nextElementSibling.innerHTML = "";
    flag1 = true;
  } else {
    flag1 = false;
    this.nextElementSibling.className = "wrong";
    this.nextElementSibling.innerHTML = "&nbsp输入格式错误";
  }
});
pwd.addEventListener("blur", function () {
  if (upwd.test(this.value)) {
    this.nextElementSibling.innerHTML = "";
    flag2 = true;
  } else {
    flag2 = false;
    this.nextElementSibling.className = "wrong";
    this.nextElementSibling.innerHTML = "&nbsp输入格式错误";
  }
});
tel.addEventListener("blur", function () {
  if (utel.test(this.value)) {
    this.nextElementSibling.innerHTML = "";
    flag3 = true;
  } else {
    flag3 = false;
    this.nextElementSibling.className = "wrong";
    this.nextElementSibling.innerHTML = "&nbsp输入格式错误";
  }
});
btn2.addEventListener("blur", function () {
  if (flag1 == true && flag2 == true && flag3 == true) {
    axios({
      method: "get",
      url: "127.0.0.1.9090/register",
    }).then(reponse => {
      console.log(reponse);
      console.log(reponse.code);
      if ((reponse.status = 1)) {
        window.location.href = "https://www.baidu.com/";
      }
    });
  } else {
    alert("请检查您输入的信息是否有误");
  }
});
