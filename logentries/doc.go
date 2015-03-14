/*
Package logentries provides a tidy log backend for the real-time
log analytics and management platform https://logentries.com/.

Use the `logentries.Configure` method to create a backend instance. The
only information that is required is your log token. This can be
acquired at http://logentries.com/.

    token := "2bfbea1e-10c3-4419-bdad-7e6435882e1f"
    backend := logentries.Configure(token).Build()
*/
package logentries
