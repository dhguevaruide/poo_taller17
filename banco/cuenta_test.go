package banco

import (
	"testing"
)

// TestDepositarYSaldo prueba que Depositar funcione y ObtenerSaldo devuelva lo correcto
func TestDepositarYSaldo(t *testing.T) {
	// 1. CREAR: Una cuenta nueva
	cuenta := NuevaCuenta("Ana Pérez")

	// 2. VERIFICAR: El saldo inicial es 0
	saldoInicial := cuenta.ObtenerSaldo()
	if saldoInicial != 0 {
		t.Errorf("Saldo inicial debería ser 0, pero es %f", saldoInicial)
	}

	// 3. DEPOSITAR: 100 pesos
	cuenta.Depositar(100)

	// 4. VERIFICAR: El saldo ahora es 100
	saldoFinal := cuenta.ObtenerSaldo()
	if saldoFinal != 100 {
		t.Errorf("Saldo después de depositar 100 debería ser 100, pero es %f", saldoFinal)
	}
}

// TestDepositarMontoNegativo prueba que no se pueda depositar dinero negativo
func TestDepositarMontoNegativo(t *testing.T) {
	cuenta := NuevaCuenta("Carlos Ruiz")

	// Intentar depositar -50 (negativo)
	cuenta.Depositar(-50)

	// El saldo debería seguir siendo 0
	saldo := cuenta.ObtenerSaldo()
	if saldo != 0 {
		t.Errorf("Depositar -50 no debería cambiar el saldo, pero es %f", saldo)
	}
}

// TestMultiplesDepositos prueba varios depósitos seguidos
func TestMultiplesDepositos(t *testing.T) {
	cuenta := NuevaCuenta("María López")

	cuenta.Depositar(100)
	cuenta.Depositar(50)
	cuenta.Depositar(25)

	saldo := cuenta.ObtenerSaldo()
	if saldo != 175 {
		t.Errorf("100 + 50 + 25 debería ser 175, pero es %f", saldo)
	}
}