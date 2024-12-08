package emailtemplates

func GenerateChangePasswordEmail(fullname string, host string, token string) string {
	return `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Password Reset</title>
  <style>
    body {
      font-family: Arial, sans-serif;
      margin: 0;
      padding: 0;
      background-color: #f8f9fa;
      display: flex;
      justify-content: center;
      align-items: center;
      min-height: 100vh;
    }
    .email-container {
      max-width: 600px;
      background: #ffffff;
      border-radius: 10px;
      box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
      text-align: center;
      padding: 20px;
      overflow: hidden;
    }
    .email-header {
      background: #dc3545;
      color: #ffffff;
      padding: 15px 0;
      font-size: 1.8em;
    }
    .email-body {
      padding: 20px;
      color: #333333;
      font-size: 1.1em;
      text-align: left;
    }
    .email-body p {
      line-height: 1.6;
    }
    .email-body a {
      display: inline-block;
      margin-top: 20px;
      padding: 10px 20px;
      background: #dc3545;
      color: #ffffff;
      text-decoration: none;
      border-radius: 5px;
      font-weight: bold;
    }
    .email-body a:hover {
      background: #bd2130;
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
      Password Reset Request
    </div>
    <div class="email-body">
      <p>Heyy ` + fullname + `,</p>
      <p>Your link to reset your Switchr account password is below:</p>
      <a href="` + host + `/changepass/` + token + `">Reset Your Password</a>
      <p>The link is active for only 5 minutes. If you didn’t request this reset, you can safely ignore this email.</p>
      <p>If the button above doesn’t work, copy and paste this link into your browser:</p>
      <p><code>` + host + `/changepass/` + token + `</code></p>
    </div>
    <div class="email-footer">
      Thank you for using Switchr!
    </div>
  </div>
</body>
</html>
`
}
