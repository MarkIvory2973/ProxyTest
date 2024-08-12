import numpy, sys
from click import command, option
from rich.console import Console
from rich.table import Table
from rich.color import Color
from rich.style import Style
from rich.text import Text
from MiHoMoAPI import *
from MiHoMoAPI import proxies

console = Console()

def normalization(x):
    return x / numpy.std(x)

@command()
@option("--host", default="127.0.0.1", help="MiHoMo API host")
@option("--port", default="9090", help="MiHoMo API port")
@option("--https", is_flag=True, help="Use HTTPS")
@option("--excludes", default="", help="Remove exclusions")
@option("--group", help="Group name")
@option("--weight", "-k", default=0.5, type=float, help="Weight (0~1)")
def main(host, port, https, excludes, group, weight):
    version = set_api_url(host, port, https)
    if version:
        console.print(f"[[green bold]*[/]] Connected to MiHoMo([deep_sky_blue2 bold]v{version}[/])")
    else:
        console.print("[[red bold]![/]] Unable to connect MiHoMo([deep_sky_blue2 bold]Unknown[/])")
        sys.exit(1)
    set_excludes([exclude for exclude in excludes.split(",") if exclude])
    
    proxies_delays = proxies.get_proxies_delays(group)
    if proxies_delays:
        console.print(f"[[green bold]*[/]] Fetched {len(proxies_delays)} delays of proxies")
    else:
        console.print(f"[[red bold]![/]] Unable to fetch the delays of proxies")
        sys.exit(1)
        
    result_length = [len(proxy_delays["delays"]) for proxy_delays in proxies_delays]
    result = numpy.array([proxy_delays["delays"][:min(result_length)] for proxy_delays in proxies_delays])
    stability = numpy.std(result, 1).reshape(len(result), 1)
    delay = numpy.percentile(result, 50, 1, keepdims=True)
    scaled_stability = normalization(stability)
    scaled_delay = normalization(delay)
    score = (1-weight)*scaled_stability + weight*scaled_delay

    table = Table(title="[deep_sky_blue2 italic]Result")
    table.add_column("Name", justify="left")
    table.add_column("Score", justify="right")
    table.add_column("Stability", justify="right")
    table.add_column("Delay", justify="right")
    idxs = numpy.argsort(score.T).tolist()[0]
    Score_ratio = ((score-min(score))/(max(score)-min(score))).T.tolist()[0]
    for idx in idxs:
        Name = proxies_delays[idx]["name"]
        Score = Text(f"{(1-score[idx].item())*100:.2f}", Style(color=Color.from_rgb(int(255*Score_ratio[idx]), int(255*(1-Score_ratio[idx])), 0)))
        Stability = f"{stability[idx].item():.2f}"
        Delay = str(int(delay[idx].item()))
        
        table.add_row(Name, Score, Stability, Delay)
    console.print()
    console.print(table)
    
if __name__ == "__main__":
    main()