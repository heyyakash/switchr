package emailtemplates

func GenerateMagicLinkEmailTemplate(host string, token string) string {
	return `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Login Link</title>
  <style>
    body {
      font-family: Arial, sans-serif;
      margin: 0;
      padding: 0;
      background-color: #f9f9f9;
      display: flex;
      justify-content: center;
      align-items: center;
      height: 100vh;
    }
    .email-container {
      max-width: 500px;
      background: #ffffff;
      border-radius: 10px;
      box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
      overflow: hidden;
      text-align: center;
      padding: 20px;
    }
    .email-header {
      background: #4CAF50;
      color: #ffffff;
      padding: 15px 0;
      font-size: 1.5em;
    }
    .email-body {
      padding: 20px;
      color: #333333;
      font-size: 1.1em;
    }
    .email-body a {
      display: inline-block;
      margin-top: 20px;
      padding: 10px 15px;
      background: #4CAF50;
      color: white;
      text-decoration: none;
      border-radius: 5px;
      font-weight: bold;
    }
    .email-body a:hover {
      background: #45a049;
    }
    .email-footer {
      margin-top: 20px;
      font-size: 0.9em;
      color: #888888;
    }
  </style>
</head>
<body>
  <div class="email-container">
    <div class="email-header">
      Login Link
    </div>
    <div class="email-body">
      <p>Heyy! Your login link is as follows and is only valid for 5 minutes:</p>
      <a href="` + host + `/user/magic/verify/` + token + `">Click here to log in</a>
    </div>
    <div class="email-footer">
      If you did not request this link, please ignore this email.
    </div>
  </div>
</body>
</html>
`
}
