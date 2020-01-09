package static

const (
	AlreadyWon = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Already Won</title>
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css"
          integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">

</head>
<body>
<div class="card">
    <div class="card-header font-weight-bolder">
        WINNER
    </div>
    <div class="card-body container-fluid">
        <h5 class="card-title">You already won</h5>
        <p class="card-text">Please check your email</p>
        <a href="https://raffleit.netlify.com" class="btn btn-primary">See Other winners</a>
    </div>
</div>
</body>
</html>
`

	Congrats = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Congratulations</title>
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css"
          integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">
</head>
<body>
<div class="card">
    <div class="card-header font-weight-bolder">
        Congratulations
    </div>
    <div class="card-body container-fluid">
        <h5 class="card-title">Hurray!</h5>
        <p class="card-text">You won. Your draw was magical. You'll receive an email from us soon.</p>
        <a href="https://raffleit.netlify.com" class="btn btn-primary">See Other winners</a>
    </div>
</div>
</body>
</html>
`

	InvalidReq = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Invalid Request</title>
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css"
          integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">

</head>
<body>
<div class="card">
    <div class="card-header font-weight-bolder">
        INVALID REQUEST
    </div>
    <div class="card-body container-fluid">
        <h5 class="card-title">Don't do that again</h5>
        <p class="card-text">Request seems to be poorly formatted or missing some query parameters</p>
        <a href="https://raffleit.netlify.com" class="btn btn-primary">Place raffle</a>
    </div>
</div>
</body>
</html>
`

	InvalidTx = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Invalid Transaction</title>
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css"
          integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">

</head>
<body>
<div class="card">
    <div class="card-header font-weight-bolder">
        TRANSACTION ERROR
    </div>
    <div class="card-body container-fluid">
        <h5 class="card-title">Couldn't verify Transaction</h5>
        <p class="card-text">Please proceed to make raffle and payment.</p>
        <a href="https://raffleit.netlify.com" class="btn btn-primary">Go to Raffle Page</a>
    </div>
</div>
</body>
</html>
`

	PaymentRef = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Used Payment Ref</title>
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css"
          integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">

</head>
<body>
<div class="card">
    <div class="card-header font-weight-bolder">
        PAYMENT ERROR
    </div>
    <div class="card-body container-fluid">
        <h5 class="card-title">REF has already been used</h5>
        <p class="card-text">Please proceed to make another raffle and payment.</p>
        <a href="https://raffleit.netlify.com" class="btn btn-primary">Go to Raffle Page</a>
    </div>
</div>
</body>
</html>`

	Sorry = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Sorry</title>
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css"
          integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">

</head>
<body>
<div class="card">
    <div class="card-header font-weight-bolder">
        SORRY
    </div>
    <div class="card-body container-fluid">
        <h5 class="card-title">Oh no!</h5>
        <p class="card-text">That didn't go as expected. You can always try again!</p>
        <a href="https://raffleit.netlify.com" class="btn btn-primary">Try again</a>
        <a href="https://raffleit.netlify.com" class="btn btn-secondary ">See winners</a>

    </div>
</div>
</body>
</html>`

	Winners = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Winners</title>
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css"
          integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">

</head>
<body>
<div class="container pt-3">
<div class="float-right">
        <a  href="https:/raffleit.netlify.com">Go Back to Raffle</a>
    </div>
    <h1>Cheers to our <span class="text-info font-weight-bolder">Winners</span></h1>
    <div class="card mt-3">
        <table class="table table-striped">
            <thead>
            <tr>
                <th scope="col">#</th>
                <th scope="col">Name</th>
                <th scope="col">Phone</th>
                <th scope="col">Email</th>
            </tr>
            </thead>
            <tbody>
            %s
            </tbody>
        </table>

    </div>

</div>
</body>
</html>`
)


