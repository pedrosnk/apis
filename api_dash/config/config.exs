# This file is responsible for configuring your application
# and its dependencies with the aid of the Mix.Config module.
#
# This configuration file is loaded before any dependency and
# is restricted to this project.
use Mix.Config

# General application configuration
config :api_dash,
  ecto_repos: [ApiDash.Repo]

# Configures the endpoint
config :api_dash, ApiDash.Endpoint,
  url: [host: "localhost"],
  secret_key_base: "IABP0mxoeAD7+1r/g1R7KI5pW/fjqB7H1lcltN6MMIzigJ7c7UuFqZodc2wMBkGV",
  render_errors: [view: ApiDash.ErrorView, accepts: ~w(html json)],
  pubsub: [name: ApiDash.PubSub,
           adapter: Phoenix.PubSub.PG2]

# Configures Elixir's Logger
config :logger, :console,
  format: "$time $metadata[$level] $message\n",
  metadata: [:request_id]

# Import environment specific config. This must remain at the bottom
# of this file so it overrides the configuration defined above.
import_config "#{Mix.env}.exs"
