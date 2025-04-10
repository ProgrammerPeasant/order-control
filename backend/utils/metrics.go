package utils

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Metrics содержит счетчики и метрики Prometheus
type Metrics struct {
	TotalRequests       *prometheus.CounterVec
	ResponseTime        *prometheus.HistogramVec
	RequestsInFlight    prometheus.Gauge
	ErrorsTotal         *prometheus.CounterVec
	DatabaseErrorsTotal *prometheus.CounterVec

	UserRegistrations     *prometheus.CounterVec
	UserRoleRegistrations *prometheus.CounterVec
	UserLogins            *prometheus.CounterVec
	AuthFailures          *prometheus.CounterVec
	ActiveUsers           prometheus.Gauge
	AuthOperationDuration *prometheus.HistogramVec

	JoinRequestOperationsTotal *prometheus.CounterVec
	JoinRequestTotal           *prometheus.GaugeVec
	JoinRequestDuration        *prometheus.HistogramVec
	PendingJoinRequestsGauge   *prometheus.GaugeVec
}

// NewMetrics создает и регистрирует метрики Prometheus
func NewMetrics() *Metrics {
	metrics := &Metrics{
		TotalRequests: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "order_control_http_requests_total",
				Help: "Общее количество HTTP запросов",
			},
			[]string{"method", "endpoint", "status"},
		),
		ResponseTime: promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "order_control_http_response_time_seconds",
				Help:    "Время ответа в секундах",
				Buckets: []float64{0.1, 0.3, 0.5, 0.7, 1, 3, 5, 10},
			},
			[]string{"method", "endpoint"},
		),
		RequestsInFlight: promauto.NewGauge(
			prometheus.GaugeOpts{
				Name: "order_control_http_requests_in_flight",
				Help: "Текущее количество запросов в обработке",
			},
		),
		ErrorsTotal: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "order_control_errors_total",
				Help: "Общее количество ошибок по типам",
			},
			[]string{"type", "message"},
		),
		DatabaseErrorsTotal: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "order_control_db_errors_total",
				Help: "Общее количество ошибок базы данных по операциям",
			},
			[]string{"operation", "table"},
		),

		// auth здесь
		UserRegistrations: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "order_control_user_registrations_total",
				Help: "Общее количество регистраций пользователей",
			},
			[]string{"registration_type"}, // "standard" или "admin"
		),
		UserRoleRegistrations: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "order_control_user_role_registrations_total",
				Help: "Количество регистраций пользователей по ролям",
			},
			[]string{"role"}, // "ADMIN", "USER", "MANAGER"
		),
		UserLogins: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "order_control_user_logins_total",
				Help: "Общее количество успешных входов в систему по ролям",
			},
			[]string{"role"},
		),
		AuthFailures: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "order_control_auth_failures_total",
				Help: "Общее количество неудачных попыток входа",
			},
			[]string{"username"},
		),
		ActiveUsers: promauto.NewGauge(
			prometheus.GaugeOpts{
				Name: "order_control_active_users",
				Help: "Примерное количество активных пользователей (увеличивается при входе)",
			},
		),
		AuthOperationDuration: promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "order_control_auth_operation_duration_seconds",
				Help:    "Время выполнения операций авторизации в секундах",
				Buckets: []float64{0.01, 0.05, 0.1, 0.5, 1, 2.5, 5},
			},
			[]string{"operation"}, // "login", "register", "admin_register"
		),

		// Метрики для запросов на присоединение
		JoinRequestOperationsTotal: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "order_control_join_request_operations_total",
				Help: "Общее количество операций с запросами на присоединение",
			},
			[]string{"operation", "result"}, // operation: create, approve, reject, get_pending; result: attempt, success, error
		),
		JoinRequestTotal: promauto.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "order_control_join_requests_total",
				Help: "Общее количество запросов на присоединение по статусам",
			},
			[]string{"status"}, // status: pending, approved, rejected
		),
		JoinRequestDuration: promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "order_control_join_request_duration_seconds",
				Help:    "Время выполнения операций с запросами на присоединение",
				Buckets: []float64{0.01, 0.05, 0.1, 0.5, 1, 2.5, 5},
			},
			[]string{"operation"}, // operation: create, approve, reject, get_pending
		),
		PendingJoinRequestsGauge: promauto.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "order_control_pending_join_requests",
				Help: "Количество ожидающих запросов на присоединение по компаниям",
			},
			[]string{"company_id"},
		),
	}

	return metrics
}

// RegisterError регистрирует ошибки приложения
func (m *Metrics) RegisterError(errorType, message string) {
	m.ErrorsTotal.WithLabelValues(errorType, message).Inc()
}

// RegisterDBError регистрирует ошибки базы данных
func (m *Metrics) RegisterDBError(operation, table string) {
	m.DatabaseErrorsTotal.WithLabelValues(operation, table).Inc()
}
