package gonet

import (
	"fmt"
	"strings"
)

const NEURON_InputTypeNeuron = 1
const NEURON_InputTypeInData = 2

type NeuronInput struct {
	InputType int
	InputID int
	Weight float64
}

type Neuron struct {
	Id int
	LayerID int
	Output float64
	Inputs []NeuronInput
}

type NeuralNet struct {
	Neurons []Neuron
	InData []float64
}

func SaveNet(filename string, network NeuralNet) (error) {
	data := []string{fmt.Sprintf("%d;%d", len(network.Neurons), len(network.InData))}
	for neuron_id, neuron := range network.Neurons {
		base_info := fmt.Sprintf("%d;%d", neuron_id, neuron.LayerID)
		input_info := []string{}
		for _, input := range neuron.Inputs {
			input_info = append(input_info,
				fmt.Sprintf("%d|%d|%f", input.InputID, input.InputType, input.Weight))
		}
		base_info += ";" + strings.Join(input_info, ";")
	}

	return WriteLines(data, filename)

}

func LoadNet(filename string) (NeuralNet) {
	//todo write the func
	return CreateEmpty([]int{}, 0)
}

func CreateEmpty(layers []int, num_inputs int) (NeuralNet) {
	result := NeuralNet{
		Neurons: []Neuron{},
		InData: make([]float64, num_inputs),
	}
	id:=0
	for layer, num_of_neurons := range layers {
		for i:=0; i<num_of_neurons; i++ {
			result = CreateNeuron(&result, Neuron{
				Id: id,
				LayerID: layer,
			})
			id++
		}
	}

	return result
}

func AppendNeuronInput(n Neuron, input NeuronInput, weightrandom bool ) (Neuron) {
	if (weightrandom) {
		input.Weight = GetFloat64(0.5, 1)
	}
	n.Inputs = append(n.Inputs, input)
	return n
}


func CreateNeuron(thenet NeuralNet, n Neuron) (NeuralNet) {
	if (n.LayerID == 0) {
		l := len(thenet.InData)
		for i:=0; i<l; i++ {
			n = AppendNeuronInput(&n, &NeuronInput{
				InputType:NEURON_InputTypeInData,
				InputID:i,
			}, true)
		}
	} else {
		l:= len(thenet.Neurons)
		for i:=l-1; i>-1; i-- {
			if (thenet.Neurons[i].LayerID +1 < n.LayerID) {
				break;
			}else if(thenet.Neurons[i].LayerID +1 > n.LayerID) {
				continue;
			}
			n = AppendNeuronInput(&n, &NeuronInput{
				InputType:NEURON_InputTypeNeuron,
				InputID:i,
			}, true)
		}
	}
	thenet.Neurons = append(thenet.Neurons, n)
	return thenet
}