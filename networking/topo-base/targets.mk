## Base topology lab
lab-base:
	sudo containerlab deploy -t topo-base/topo.yml --reconfigure

## Base topology cleanup
	sudo containerlab -t topo-base/topo.yml --cleanup