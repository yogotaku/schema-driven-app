const hooks = require("hooks");

hooks.before(
  "/users/{userId} > Get User Info by User ID > 400 > application/json",
  (transaction) => {
    const regex = /\/users\/[0-9]+/;
    const param = "/users/0";
    transaction.fullPath = transaction.fullPath.replace(regex, param);
  }
);

hooks.before(
  "/users/{userId} > Update User Information > 400 > application/json",
  (transaction) => {
    const requestBody = JSON.parse(transaction.request.body);
    requestBody["firstName"] = "";
    transaction.request.body = JSON.stringify(requestBody);
  }
);

hooks.before(
  "/users > Create New User > 400 > application/json",
  (transaction) => {
    const requestBody = JSON.parse(transaction.request.body);
    requestBody["firstName"] = "";
    transaction.request.body = JSON.stringify(requestBody);
  }
);
