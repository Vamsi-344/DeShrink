import * as React from 'react'
import { TanStackRouterDevtools } from "@tanstack/react-router-devtools"
import { Outlet, createRootRoute } from '@tanstack/react-router'

export const Route = createRootRoute({
  component: RootComponent,
})

function RootComponent() {
  return (
    <React.Fragment>
      <div>Hello "__root"!</div>
      <Outlet />
      <TanStackRouterDevtools/>
    </React.Fragment>
  )
}
