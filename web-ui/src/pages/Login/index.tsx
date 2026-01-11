import { Button } from '@/components/ui/button';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { useNavigate } from 'react-router';

export default function Login() {
  const navigate = useNavigate();
  return (
    <div className="bg-background flex flex-col overflow-hidden rounded-lg border bg-clip-padding md:flex-1 xl:rounded-xl">
      <div className="relative container hidden min-h-screen flex-1 shrink-0 items-center justify-center md:grid lg:max-w-none lg:grid-cols-2 lg:px-0">
        <div className="relative hidden h-full flex-col bg-linear-to-br from-primary/10 via-primary/5 to-background p-10 lg:flex dark:border-r">
          <div className="absolute inset-0 bg-linear-to-br from-primary/20 via-transparent to-primary/10"></div>
          <div className="relative z-20 flex items-center text-xl font-bold">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              strokeWidth="2"
              strokeLinecap="round"
              strokeLinejoin="round"
              className="mr-3 h-8 w-8 text-primary"
            >
              <path d="M15 6v12a3 3 0 1 0 3-3H6a3 3 0 1 0 3 3V6a3 3 0 1 0-3 3h12a3 3 0 1 0-3-3"></path>
            </svg>
            <span className="bg-linear-to-r from-primary to-primary/80 bg-clip-text text-transparent">Template</span>
          </div>
          <div className="relative z-20 flex-1 flex flex-col justify-center">
            <div className="space-y-8">
              {/* Financial Chart Illustration */}
              <div className="relative">
                <img
                  src="https://img.keaitupian.cn/newupload/05/1716791283753363.jpg"
                  alt="dashboard preview"
                  className="w-full h-72 object-cover rounded-lg opacity-80 mix-blend-multiply dark:mix-blend-screen"
                />
              </div>

              <div className="space-y-4">
                <h1 className="text-3xl font-bold tracking-tight text-foreground">Welcome back</h1>
                <p className="text-lg text-muted-foreground leading-relaxed">
                  Sign in to access your trading dashboard and manage your portfolio with confidence.
                </p>
              </div>

              <div className="flex items-center space-x-2 text-sm text-muted-foreground">
                <svg className="h-4 w-4" fill="currentColor" viewBox="0 0 20 20">
                  <path fillRule="evenodd" d="M10 1L13 7h7l-5.5 4L17 19l-7-5-7 5 2.5-8L0 7h7l3-6z" clipRule="evenodd" />
                </svg>
                <span>Secure • Fast • Reliable</span>
              </div>
            </div>
          </div>
        </div>
        <div className="flex items-center justify-center lg:p-8">
          <div className="mx-auto flex w-full flex-col justify-center gap-8 sm:w-95">
            <div className="text-center space-y-2">
              <div className="inline-flex items-center justify-center w-16 h-16 rounded-full bg-primary/10 mb-4">
                <svg className="w-8 h-8 text-primary" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"
                  />
                </svg>
              </div>
              <h1 className="text-3xl font-bold tracking-tight">Welcome back</h1>
              <p className="text-muted-foreground text-lg">Sign in to your TradeFlow account</p>
            </div>

            <Card className="border-0 shadow-lg bg-card/50 backdrop-blur-sm">
              <CardHeader className="space-y-1 pb-4">
                <CardTitle className="text-xl text-center">Sign in</CardTitle>
                <CardDescription className="text-center">Choose your preferred sign-in method</CardDescription>
              </CardHeader>
              <CardContent className="space-y-4">
                <Button
                  variant="outline"
                  className="w-full h-12 text-base font-medium border-2 hover:bg-primary hover:text-primary-foreground transition-all duration-200 hover:scale-[1.02] hover:shadow-md"
                  onClick={() => {
                    sessionStorage.setItem('logged', 'true');
                    navigate('/');
                  }}
                >
                  <svg className="mr-3 h-5 w-5" viewBox="0 0 24 24">
                    <path
                      fill="currentColor"
                      d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z"
                    />
                    <path
                      fill="currentColor"
                      d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z"
                    />
                    <path
                      fill="currentColor"
                      d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z"
                    />
                    <path
                      fill="currentColor"
                      d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z"
                    />
                  </svg>
                  Continue with Google
                </Button>

                <div className="relative">
                  <div className="absolute inset-0 flex items-center">
                    <div className="w-full border-t border-border/50"></div>
                  </div>
                  <div className="relative flex justify-center text-xs uppercase">
                    <span className="bg-background px-3 text-muted-foreground font-medium">Secure & Encrypted</span>
                  </div>
                </div>

                <div className="flex items-center justify-center space-x-4 text-xs text-muted-foreground">
                  <div className="flex items-center space-x-1">
                    <svg className="w-3 h-3 text-green-500" fill="currentColor" viewBox="0 0 20 20">
                      <path
                        fillRule="evenodd"
                        d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
                        clipRule="evenodd"
                      />
                    </svg>
                    <span>SSL Protected</span>
                  </div>
                  <div className="flex items-center space-x-1">
                    <svg className="w-3 h-3 text-blue-500" fill="currentColor" viewBox="0 0 20 20">
                      <path
                        fillRule="evenodd"
                        d="M5 9V7a5 5 0 0110 0v2a2 2 0 012 2v5a2 2 0 01-2 2H5a2 2 0 01-2-2v-5a2 2 0 012-2zm8-2v2H7V7a3 3 0 016 0z"
                        clipRule="evenodd"
                      />
                    </svg>
                    <span>256-bit Encryption</span>
                  </div>
                </div>
              </CardContent>
            </Card>

            <div className="text-center space-y-2">
              <p className="text-sm text-muted-foreground">
                By signing in, you agree to our{' '}
                <a href="/terms" className="text-primary hover:underline font-medium">
                  Terms of Service
                </a>{' '}
                and{' '}
                <a href="/privacy" className="text-primary hover:underline font-medium">
                  Privacy Policy
                </a>
              </p>
              <p className="text-xs text-muted-foreground">
                Need help?{' '}
                <a href="/support" className="text-primary hover:underline">
                  Contact Support
                </a>
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
