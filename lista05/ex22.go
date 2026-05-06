// 22. Faça um programa que simule um controle bancário. Para tanto, devem ser lidos os códigos de dez contas e os
// seus respectivos saldos. Os códigos devem ser armazenados em um vetor de números inteiros (não pode haver
// mais do que uma conta com o mesmo código) e os saldos devem ser armazenados em um vetor de números
// reais. O saldo deverá ser cadastrado na mesma posição do código. Por exemplo, se a conta 504 foi armazenada
// na 5a posição do vetor de códigos, o seu saldo deverá ficar na 5a posição do vetor de saldos. Depois de fazer a
// leitura dos valores, mostrar o seguinte menu na tela:
// 1. Efetuar depósito
// 2. Efetuar saque
// 3. Consultar o ativo bancário (ou seja, o somatório dos saldos de todos os clientes)
// 4. Finalizar o programa.
// - Para efetuar depósito deve-se solicitar o código da conta e o valor a ser depositado. Se a conta não estiver
// cadastrada, mostrar a mensagem Conta não encontrada e voltar ao menu. Se a conta existir, atualizar o seus
// saldo.
// - Para efetuar saque deve-se solicitar o código da conta e o valor a ser sacado. Se a conta não estiver
// cadastrada, mostrar a mensagem Conta não encontrada e voltar ao menu. Se a conta existir, verificar verificar
// se o seu saldo é suficiente para cobrir o saque. (Estamos supondo que a conta não pode ficar com o saldo
// negativo). Se o saldo for suficiente, realizar o saque e voltar ao menu. Caso contrário, mostrar a mensagem
// Saldo insuficiente e voltar ao menu.
// - Para consultar o ativo bancário deve-se somar o saldo de todas as contas do banco. Depois de mostrar esse
// valor, voltar ao menu.
// - O programa só termina quando for digitada a opção 4 – Finalizar programa.