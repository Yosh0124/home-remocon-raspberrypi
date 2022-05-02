export const ControlAir: (action: string) => Promise<Response> = async (action: string) => {
  const apiHost: string = 'http://localhost:1323'

  const headers: { [name: string]: string } = {
    'Content-Type': 'application/json',
    'Accept': 'application/json',
  }

  const postAction: {[name: string]: string} = {'status': action}

  return fetch(
    apiHost + '/air',
    {
      body: JSON.stringify(postAction),
      headers: headers,
      method: 'POST',
    }
  )
}