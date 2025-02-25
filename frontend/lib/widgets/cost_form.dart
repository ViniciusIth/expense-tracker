import 'package:everything_alright/services/google_api.dart';
import 'package:flutter/material.dart';

class CostForm extends StatefulWidget {
  const CostForm({super.key});

  @override
  State<StatefulWidget> createState() => CostFormState();
}

class CostFormState extends State<CostForm> {
  final _formKey = GlobalKey<FormState>();
  final _costController = TextEditingController();
  final _descriptionController = TextEditingController();
  final _observationController = TextEditingController();

  final List<String> _categories = [
    "Transporte",
    "Saúde",
    "Alimentação",
    "Recorrentes",
    "Recreação",
    "Pessoal",
  ];
  String? _selectedCategory;

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(16.0),
      child: Form(
        key: _formKey,
        child: SingleChildScrollView(
          child: Column(
            children: [
              _buildDescriptionField(),
              const SizedBox(height: 20),
              _buildCategoryDropdownField(),
              const SizedBox(height: 20),
              _buildCostField(),
              const SizedBox(height: 20),
              _buildObservationField(),
              const SizedBox(height: 20),
              _buildSubmitButton(),
            ],
          ),
        ),
      ),
    );
  }

  Widget _buildDescriptionField() {
    return TextFormField(
      decoration: const InputDecoration(
        labelText: 'Com o que gastei',
        border: OutlineInputBorder(),
      ),
      validator: (value) {
        if (value == null || value.isEmpty) {
          return 'Insira o motivo do gasto';
        }
        return null;
      },
      controller: _descriptionController,
    );
  }

  Widget _buildCategoryDropdownField() {
    return DropdownButtonFormField(
      value: _selectedCategory,
      items:
          _categories.map((String category) {
            return DropdownMenuItem<String>(
              value: category,
              child: Text(category),
            );
          }).toList(),
      onChanged: (String? newValue) {
        setState(() {
          _selectedCategory = newValue;
        });
      },
      validator: (value) {
        if (value == null || value.isEmpty) {
          return 'Please select a category';
        }
        return null;
      },
      decoration: const InputDecoration(
        labelText: 'Category',
        border: OutlineInputBorder(),
      ),
    );
  }

  Widget _buildCostField() {
    return TextFormField(
      controller: _costController,
      keyboardType: TextInputType.numberWithOptions(decimal: true),
      validator: (value) {
        if (value == null || value.isEmpty) {
          return 'Please enter the amount';
        }
        if (double.tryParse(value) == null) {
          return 'Please enter a valid number';
        }
        return null;
      },
      decoration: const InputDecoration(
        labelText: 'Amount',
        border: OutlineInputBorder(),
        prefixText: 'R\$',
      ),
    );
  }

  Widget _buildObservationField() {
    return TextFormField(
      decoration: const InputDecoration(
        labelText: 'Observações',
        border: OutlineInputBorder(),
      ),
      controller: _observationController,
    );
  }

  Widget _buildSubmitButton() {
    return ElevatedButton(
      onPressed: () async {
        if (_formKey.currentState!.validate()) {
          final category = _selectedCategory!;
          final amount = double.tryParse(_costController.text);
          final description = _descriptionController.text;
          final observation =
              _observationController.text.isNotEmpty
                  ? _observationController.text
                  : '';

          final response = await appendToSheet(
            rowData: [description, category, amount, observation],
          );

          if (!mounted) return;

          if (response['status'] == 'SUCCESS') {
            ScaffoldMessenger.of(context).showSnackBar(
              const SnackBar(content: Text('Expense saved to Google Sheets!')),
            );

            // Clear the form
            _formKey.currentState!.reset();
            setState(() {
              _selectedCategory = null;
            });
            _costController.clear();

            return;
          }

          ScaffoldMessenger.of(context).showSnackBar(
            SnackBar(
              content: Text('Failed to save expense: ${response["message"]}'),
            ),
          );
        }
      },
      child: const Text('Save Expense'),
    );
  }
}
