package emailtemplates

func GenerateVerificationEmail(fullname string, host string, token string) string {
	return `
	<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Email Verification</title>
  <style>
    body {
      font-family: Arial, sans-serif;
      margin: 0;
      padding: 0;
      background-color: #f4f4f4;
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
      background: #007bff;
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
      background: #007bff;
      color: #ffffff;
      text-decoration: none;
      border-radius: 5px;
      font-weight: bold;
    }
    .email-body a:hover {
      background: #0056b3;
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
      Email Verification
    </div>
    <div class="email-body">
      <p>Dear ` + fullname + `,</p>
      <p>We hope this message finds you well. To complete your registration process, please verify your email address by clicking the button below:</p>
      <a href="` + host + `/user/verify/` + token + `">Verify Your Email</a>
      <p>If the button above doesn't work, copy and paste the following link into your browser:</p>
      <p><code>` + host + `/user/verify/` + token + `</code></p>
    </div>
    <div class="email-footer">
      If you did not request this email, please disregard this message.
    </div>
  </div>
</body>
</html>
`
}
